/*
 * exbuffer.c
 * yoyo 2013 https://github.com/play175/exbuffer.c
 * new BSD Licensed
 */
#ifndef ASSIST_DATA_EXBUFFER_H
#define ASSIST_DATA_EXBUFFER_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#include "./codec.h"
#include "./endian.h"

//#ifdef __cplusplus
//extern "C"
//{
//#endif

#define EXTEND_BYTES 1024

typedef struct exbuffer_value
{
	//unsigned char headLen;//
    //enum exbuffer_endian endian;//
    size_t readOffset;
    size_t putOffset;
    //size_t dlen;//
	unsigned char* buffer;
	size_t bufferlen;
	size_t packetLen;
	unsigned char *packet;

    // unsigned char *headBytes;
	// union HeadBytesS
    // {
    //     unsigned char bytes[2];
    //     unsigned short val;
    // } headS;

    // union HeadBytesL
    // {
    //     unsigned char bytes[4];
    //     unsigned long val;
    // } headL;

	//void* context;//
	//void (*recvHandle)(unsigned char*, size_t, void*);//

	exbuffer_codec* codec;
} exbuffer_t;

//////////////////////////////////////////////////////////////////////////////////////////
static exbuffer_t* exbuffer_new();

static void exbuffer_reset(exbuffer_t* value);

static void exbuffer_free(exbuffer_t** value);

static void exbuffer_printHex(unsigned char* bytes,unsigned short len);

static void exbuffer_dump(exbuffer_t* value,unsigned short len);

static size_t exbuffer_getLen(exbuffer_t* value);

static void exbuffer_put(exbuffer_t* value, unsigned char* buffer,size_t offset,size_t len);

//////////////////////////////////////////////////////////////////////////////////////////
static exbuffer_t* exbuffer_new()
{
	exbuffer_t* value = (exbuffer_t *)malloc(sizeof (exbuffer_t));

	//value->headLen = 2;
	//value->endian = EXBUFFER_BIG_ENDIAN;

	value->bufferlen = EXTEND_BYTES;
	value->buffer = (unsigned char*)malloc(value->bufferlen);
	memset(value->buffer, 0, value->bufferlen);

	value->packetLen = EXTEND_BYTES;
	value->packet = (unsigned char *)malloc(value->packetLen);
	memset(value->packet, 0, value->packetLen);

	//value->headBytes = (unsigned char *)malloc(4);
	//memset(value->headBytes, 0, 4);

	value->readOffset = 0;
	value->putOffset = 0;
	//value->dlen = 0;
	//value->recvHandle = NULL;

	return value;
}

static void exbuffer_reset(exbuffer_t* value)
{
	value->readOffset = 0;
	value->putOffset = 0;
	//value->dlen = 0;
}

static void exbuffer_free(exbuffer_t** value)
{
	free ((*value)->packet);
	(*value)->packet = NULL;

	free ((*value)->buffer);
	(*value)->buffer = NULL;

	// free ((*value)->headBytes);
	// (*value)->headBytes = NULL;

	// (*value)->recvHandle = NULL;

	// (*value)->context = NULL;

	free (*value);
	(*value) = NULL;
}

static void exbuffer_printHex(unsigned char* bytes,unsigned short len)
{
    if(len>50)len=50;
	unsigned short iLoop;
	for(iLoop = 0;iLoop < len;iLoop++)
    {
        printf("%02x ",bytes[iLoop]);
    }
	printf("\n");
}

static void exbuffer_dump(exbuffer_t* value,unsigned short len)
{
	exbuffer_printHex(value->buffer,len);
}

static size_t exbuffer_getLen(exbuffer_t* value)
{
	return value->putOffset - value->readOffset;
}

static void exbuffer_proc(exbuffer_t* value)
{
	size_t i;
	unsigned char rlen = 0;

	while(1)
	{
		// 查看是否解析到数据长度
		if(value->codec->GetDataLen() == 0){
			// 先判断长度有没有超过头长度
			unsigned int headerLen = value->codec->GetHeaderLen();
			if(exbuffer_getLen(value) < headerLen){
				break;
			}
			// 处理头部，并获取数据长度
			int ok = value->codec->DecodeHeader(value->buffer, value->readOffset);
			if (ok != 0){
				break;
			}

			// 移动读取
			value->readOffset = headerLen;
		}

		unsigned int dataLen = value->codec->GetDataLen();
		if(dataLen!=0 && exbuffer_getLen(value)>=dataLen){
			// 数据已经完整
			if(value->packetLen < dataLen)
			{
				size_t rn1 = dataLen/EXTEND_BYTES;
				if(dataLen%EXTEND_BYTES>0)rn1+=1;
				size_t expacketLen = rn1 * EXTEND_BYTES;

				value->packetLen = expacketLen;
				value->packet = (unsigned char *)realloc(value->packet,value->packetLen);
			}

			memcpy(value->packet,value->buffer + value->readOffset, dataLen);
			// if(value->recvHandle != NULL)
			// {
			// 	value->recvHandle(value->packet, dataLen, value->context);
			// }
			value->codec->DecodeData(value->packet, dataLen);

			value->readOffset += dataLen;

			memmove(value->buffer, value->buffer + value->readOffset, value->bufferlen-value->readOffset);
			value->putOffset -= value->readOffset;
			value->readOffset = 0;
			//value->dlen = 0;
		}else{
			break;
		}
	}
}

static void exbuffer_put(exbuffer_t* value, unsigned char* buffer,size_t offset,size_t len)
{
	if((value->putOffset + len) > value->bufferlen){

		size_t rn1 = (len + value->putOffset) / EXTEND_BYTES;
		if((len + value->putOffset)%EXTEND_BYTES > 0)
			rn1 += 1;
		size_t exbufferlen = rn1 * EXTEND_BYTES;

		value->bufferlen = exbufferlen;
		value->buffer = (unsigned char*)realloc(value->buffer,value->bufferlen);

	}

	memcpy(value->buffer + value->putOffset, buffer, len);
	value->putOffset += len;

	exbuffer_proc(value);
}

//#ifdef __cplusplus
//}
//#endif

#endif // ASSIST_DATA_EXBUFFER_H
