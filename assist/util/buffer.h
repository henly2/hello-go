/** 
 * @file buffer.h
 * @brief 序列化
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 序列化 
 */
#ifndef ASSIST_UTIL_BUFFER_H
#define ASSIST_UTIL_BUFFER_H

#include <stdio.h>

#include <string>
#include <vector>

#include "endian.h"

namespace buffer 
{
//#define DEBUG_LOG
#define DEFAULT_BUFFER_SIZE 100

class Data {
    public:
        Data(bool is_endian_little = true):
            is_endian_little_(is_endian_little),
            databuffer_(NULL),
            readcursor_(0),
            writecursor_(0),
            cap_(0){
            initByCap(DEFAULT_BUFFER_SIZE);
        }
        Data(const Data &stream){
            this->initByOther(stream);
        }
        ~Data(){
            clear();
        }

        Data& operator =(const Data &stream){
            if (this != &stream){
                this->initByOther(stream);
            }
            return *this;
        }

        const char* DataPtr()const{
            return this->databuffer_;
        }
        int DataLength()const{
            return this->writecursor_;
        }

        void MoveWriteCursorBy(int offset, int needLen = 0){
            this->writecursor_ += offset;
            this->checkCap(needLen);
        }

        void MoveWriteCursorAt(int cursor, int needLen = 0){
            this->writecursor_ = cursor;
            this->checkCap(needLen);
        }

        void MoveReadCursorBy(int offset){
            this->readcursor_ += offset;
            #ifdef DEBUG_LOG
                printf("readcursor=%d\n", this->readcursor_);
            #endif
        }
        void MoveReadCursorAt(int cursor){
            this->readcursor_ = cursor;
            #ifdef DEBUG_LOG
                printf("readcursor=%d\n", this->readcursor_);
            #endif
        }

        void Dump()const{
            printf("cap=%d, length=%d\n", this->cap_, this->writecursor_);
            printf("data:\n");
            for (int i = 0; i < this->writecursor_; i++){
                if(i % 16 == 0){
                    printf("\n");
                }
                printf("%x ", this->databuffer_[i]);
            }
            printf("\ndata end\n");
        }

        void Reset(){
            writecursor_ = 0;
            readcursor_ = 0;
        }

    public:
        // 追加写入
        void writeData(const char *data, int len){
            if (data == NULL || len == 0){
                return;
            }
            this->MoveWriteCursorBy(0, len);

            #ifdef DEBUG_LOG
                printf("=====>>\n");
                printf("write cursor=%d len=%d\n", this->writecursor_, len);
            #endif

            memcpy(this->databuffer_+this->writecursor_, data, len);
            this->writecursor_ += len;

            #ifdef DEBUG_LOG
                Dump();
                printf("<<=====\n");
            #endif
        }

        Data& operator<<(char val){
            writeVal(val);
            return *this;
        }
        void writeVal(char val){
            this->writeData((const char*)&val, sizeof(char));
        } 

        Data& operator<<(unsigned char val){
            writeVal(val);
            return *this;
        }
        void writeVal(unsigned char val){
            this->writeData((const char*)&val, sizeof(char));
        } 

        Data& operator<<(unsigned short val){
            writeVal(val);
            return *this;
        }
        void writeVal(unsigned short val){
            unsigned short v = util::exchange(val, this->is_endian_little_);
            this->writeData((const char*)&v, sizeof(unsigned short));
        }

        Data& operator<<(unsigned int val){
            writeVal(val);
            return *this;
        }
        void writeVal(unsigned int val){
            unsigned int v = util::exchange(val, this->is_endian_little_);
            this->writeData((const char*)&v, sizeof(unsigned int));
        }

        Data& operator<<(const std::string& val){
            writeVal(val);
            return *this;
        }
        void writeVal(const std::string& val, int reservelen=-1){
            this->writeData((const char*)val.c_str(), val.size());
            if(reservelen != -1){
                this->MoveWriteCursorBy(reservelen-val.size(), 0);
            }  
        }

    public:
        bool operator>>(char &val){
            return readVal(val);
        }
        bool readVal(char &val){
            if (this->readcursor_ + sizeof(char) > this->writecursor_){
                return false;
            }

            val = *(char*)(this->databuffer_+this->readcursor_);
            this->readcursor_ += sizeof(char);
            return true;
        }

        bool operator>>(unsigned char &val){
            return readVal(val);
        }
        bool readVal(unsigned char &val){
            if (this->readcursor_ + sizeof(char) > this->writecursor_){
                return false;
            }

            val = *(unsigned char*)(this->databuffer_+this->readcursor_);
            this->readcursor_ += sizeof(char);
            return true;
        }

        bool operator>>(unsigned short &val){
            return readVal(val);
        }
        bool readVal(unsigned short &val){
            if (this->readcursor_ + sizeof(unsigned short) > this->writecursor_){
                return false;
            }

            unsigned short v = *(unsigned short*)(this->databuffer_+this->readcursor_);
            val = util::exchange(v, this->is_endian_little_);
            this->readcursor_ += sizeof(unsigned short);
            return true;
        }

        bool operator>>(unsigned int &val){
            return readVal(val);
        }
        bool readVal(unsigned int &val){
            if (this->readcursor_ + sizeof(unsigned int) > this->writecursor_){
                return false;
            }

            unsigned int v = *(unsigned int*)(this->databuffer_+this->readcursor_);
            val = util::exchange(v, this->is_endian_little_);
            this->readcursor_ += sizeof(unsigned int);
            return true;
        }

        bool operator>>(std::string &val){
            return readVal(val);
        }
        bool readVal(std::string &val, int reservelen=-1){
            if (this->readcursor_ > this->writecursor_){
                return false;
            }

            for(int i = this->readcursor_; i < this->writecursor_; i++){
                char c = *(char*)(this->databuffer_+i);
                if (c == '\0') {
                    break;
                }
                if (reservelen != -1 && i >= (this->readcursor_+reservelen)) {
                    break;
                }

                val.push_back(c);
            }
            if(reservelen != -1){
                this->readcursor_ += reservelen;
            } else {
                this->readcursor_ += val.size();
            }
            return true;
        }

    private:
        void checkCap(int needLen){
            int newlen = this->writecursor_ + needLen - this->cap_;
            if(newlen > 0){
                int cnt = newlen / DEFAULT_BUFFER_SIZE;
                this->initByCap(this->cap_ + DEFAULT_BUFFER_SIZE*(1+cnt));
            }
        }
        void initByCap(int cap){
            if (this->cap_ == cap){
                return;
            }

            char* newdatabuffer = new char[cap];
            int newcap = cap;
            int newcursor = 0;
            memset(newdatabuffer, 0, newcap);
            if (this->databuffer_ != nullptr){
                // need copy old and release old buffer
                memcpy(newdatabuffer, this->databuffer_, this->cap_);
                delete this->databuffer_;
                this->databuffer_ = nullptr;

                newcursor = this->writecursor_;
            }

            this->databuffer_ = newdatabuffer;
            this->cap_ = newcap;
            this->writecursor_ = newcursor;
        }
        void initByOther(const Data &stream){
            clear();

            this->databuffer_ = new char[stream.cap_];
            this->cap_ = stream.cap_;

            this->writecursor_ = stream.writecursor_;
            memcpy(this->databuffer_, stream.databuffer_, this->cap_);
        }
        void clear(){
            if(this->databuffer_ != NULL){
                delete this->databuffer_;
            }

            this->readcursor_ = 0;
            this->writecursor_ = 0;
            this->cap_ = 0;
        }

    private:
        bool is_endian_little_;
        char* databuffer_;
        int writecursor_;
        int readcursor_;
        int cap_;
};

}

#endif