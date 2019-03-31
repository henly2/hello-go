/** 
 * @file codec.h
 * @brief 数据解析
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 数据解析 
 */
#ifndef ASSIST_DATA_CODEC_H
#define ASSIST_DATA_CODEC_H

#include "../util/endian.h"

struct exbuffer_codec
{
    virtual unsigned int GetHeaderLen() = 0;
    virtual unsigned int GetDataLen() = 0;
    virtual int DecodeHeader(unsigned char* buffer, int cursor) = 0;
    virtual int DecodeData(unsigned char* buffer, int len) = 0;
};

struct default_exbuffer_codec : public exbuffer_codec
{
	default_exbuffer_codec()
	:dataLen(0),
	headLen(0),
	context(NULL){

	}
	virtual unsigned int GetHeaderLen(){
		return headLen;
	}
	virtual unsigned int GetDataLen(){
		return dataLen;
	}
	virtual int DecodeHeader(unsigned char* buffer, int cursor)
	{
		unsigned int i = 0;
		unsigned char headBytes[4];
		for(i=0; i<this->headLen; i++)
		{
			headBytes[i] = buffer[cursor+i];
		}
		if(this->headLen==2)
		{
			unsigned short sv = *(unsigned short*)&headBytes;
			this->dataLen = util::exchange(sv,this->endian==util::EXBUFFER_LITTLE_ENDIAN);
		}else
		{
			unsigned long lv = *(unsigned long*)&headBytes;
			this->dataLen = util::exchange(lv,this->endian==util::EXBUFFER_LITTLE_ENDIAN);
		}
		return 0;
	}
	virtual int DecodeData(unsigned char* buffer, int len)
	{
		this->recvHandle(buffer, len, this->context);
		this->dataLen = 0;
		return 0;
	}

	enum util::endian_type endian;
	unsigned int headLen;

	void* context;
	void (*recvHandle)(unsigned char*, size_t, void*);

private:
	unsigned int dataLen;
};

#endif