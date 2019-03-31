#include <stdio.h>

#include <string>

#include "./util/endian.h"
#include "./util/buffer.h"
#include "./util/exbuffer.h"

#include "./data/err.h"
#include "./data/cmd.h"
#include "./data/exdata.h"
#include "./data/excodec.h"

void testBufferData();
void testExData();
void testExBuffer();
void testExBuffer2();
int main()
{
    util::init_host_endian();
    printf("is_host_endian_little=%d\n", util::is_host_endian_little());

    //testBufferData();
    //testExData();
    //testExBuffer();
    testExBuffer2();
}

// 静态回调
static void recvHandleData(unsigned char *rbuf, size_t len, void* context)
{
    std::string out;
    out.assign((const char*)rbuf, len);

    printf("out=%s\n", out.c_str());
    
}
void testExBuffer()
{
    // 解析数据
	exbuffer_t* value_;

    // 协议
    value_ = exbuffer_new();
    // value_->headLen = 4;
    // value_->endian = EXBUFFER_LITTLE_ENDIAN;
    // value_->context = nullptr;
    // value_->recvHandle = recvHandleData;

    // codec
    default_exbuffer_codec codec;
    codec.headLen = 4;
    codec.endian = util::EXBUFFER_LITTLE_ENDIAN;
    codec.context = NULL;
    codec.recvHandle = recvHandleData;

    value_->codec = &codec;
    
    // put data
    {
        std::string in = "abcd";
        unsigned int len = in.length();
        unsigned int l = util::exchange(len, true);
        exbuffer_put(value_,(unsigned char*)&l,0,4);
        exbuffer_put(value_,(unsigned char*)in.c_str(),0,len);
    }

    // reset
    //exbuffer_reset(value_);

    // release
    if(value_!=NULL){
        exbuffer_free(&value_);
        value_ = NULL;
    }
}


// 静态回调
static void recvHandleData2(DataHeader& header, buffer::Data& cache, void* context)
{
    //std::string out;
    //out.assign((const char*)rbuf, len);

    //printf("h-l=%d\n", header.length);
    printf("h-r=%d\n", header.reqid);
    //printf("h-c=%d\n", header.cmdid);
    //cache.Dump();
}
void testExBuffer2()
{
    // 解析数据
	exbuffer_t* value_;

    // 协议
    value_ = exbuffer_new();
    // value_->headLen = 4;
    // value_->endian = EXBUFFER_LITTLE_ENDIAN;
    // value_->context = nullptr;
    // value_->recvHandle = recvHandleData;

    // codec
    ex_exbuffer_codec codec;
    codec.endian = util::EXBUFFER_LITTLE_ENDIAN;
    codec.context = NULL;
    codec.recvHandle = recvHandleData2;

    value_->codec = &codec;
    
    // put data
    for(int i= 0; i < 100; i++)
    {
        buffer::Data d;

        Req_GroupAdd req;
        req.groupInfo.name = "one";
        req.groupInfo.comment.append(64, 'c');

        DataHeader header;
        header.cmdid = Command_GroupAdd;
        header.reqid = i+1;
        header.length = req.length();

        int ret = WriteRequest(header, &req, d);
        if(ret != ECI_Ok){
            const char* msg = Get_Err_Code_Internal_Msg(ret);
            printf("err=%d, msg=%s\n", ret, msg);
            return;
        }
        //d.Dump();

        exbuffer_put(value_,(unsigned char*)d.DataPtr(),0,d.DataLength());
    }

    // reset
    //exbuffer_reset(value_);

    // release
    if(value_!=NULL){
        exbuffer_free(&value_);
        value_ = NULL;
    }
}

void testExData()
{
    buffer::Data d;

    {
        printf("write...\n");
        d.Dump();

        Req_GroupAdd req;
        req.groupInfo.name = "one";
        req.groupInfo.comment.append(64, 'c');

        DataHeader header;
        header.cmdid = Command_GroupAdd;
        header.reqid = 1;
        header.length = req.length();

        int ret = WriteRequest(header, &req, d);
        if(ret != ECI_Ok){
            const char* msg = Get_Err_Code_Internal_Msg(ret);
            printf("err=%d, msg=%s\n", ret, msg);
            return;
        }

        d.Dump();
        printf("name=%s, comment=%s(%ld-%ld)\n", req.groupInfo.name.c_str(), req.groupInfo.comment.c_str(), req.groupInfo.name.size(), req.groupInfo.comment.size());
    }

    {
        printf("read...\n");

        DataHeader header;
        int ret = ReadResponseHeader(d, header);
        if(ret != ECI_Ok){
            const char* msg = Get_Err_Code_Internal_Msg(ret);
            printf("err=%d, msg=%s\n", ret, msg);
            return;
        }

        printf("cmdid=%d\n", header.cmdid);
        if (header.cmdid == Command_GroupAdd){
            Req_GroupAdd res;
            int ret = ReadResponseData(d, header, &res);
            if(ret != ECI_Ok){
                const char* msg = Get_Err_Code_Internal_Msg(ret);
                printf("err=%d, msg=%s\n", ret, msg);
                return;
            }

            printf("name=%s, comment=%s(%ld-%ld)\n", res.groupInfo.name.c_str(), res.groupInfo.comment.c_str(), res.groupInfo.name.size(), res.groupInfo.comment.size());
        }
    }
    
}

void testBufferData()
{
    unsigned int ll = 0x12345678;
    unsigned int ll2 = util::exchange(ll, true);
    unsigned int ll3 = util::exchange(ll, false);

    printf("ll=%x, ll2=%x, ll3=%x\n", ll, ll2, ll3);

    buffer::Data d;
    d << ll;
    d.Dump();

    d << std::string("name");
    d.Dump();

    d.MoveWriteCursorBy(0, 7);
    d << std::string("comment");
    d.Dump();

    unsigned int ii = 0;
    d >> ii;
    printf("ii=%x\n", ii);

    std::string name;
    d.readVal(name, 4);
    printf("name=%s\n", name.c_str());

    std::string comment;
    d >> comment;
    printf("comment=%s\n", comment.c_str());
}