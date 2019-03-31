/** 
 * @file excodec.h
 * @brief 数据解析
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 数据解析 
 */
#ifndef ASSIST_DATA_EXCODEC_H
#define ASSIST_DATA_EXCODEC_H

#include "../util/endian.h"
#include "../util/codec.h"
#include "../util/buffer.h"

#include "./exdata.h"

struct ex_exbuffer_codec : public exbuffer_codec
{
	ex_exbuffer_codec()
	:dataLen(0),
	headLen(0),
	context(NULL){

	}
	virtual unsigned int GetHeaderLen(){
		// HeaderVersion(4) + length(4)(n+2) + reqid(4) + cmdid(4) + reqdata(n) + checksum(2) 
		return 4 + 4 + 4 + 4;
	}
	virtual unsigned int GetDataLen(){
		return dataLen;
	}
	virtual int DecodeHeader(unsigned char* buffer, int cursor)
	{
		cache.writeData((const char*)buffer, GetHeaderLen());
		
        int ret = ReadResponseHeader(cache, header);
        if(ret != ECI_Ok){
            return ret;
        }

		dataLen = header.length;
		return 0;
	}
	virtual int DecodeData(unsigned char* buffer, int len)
	{
		cache.writeData((const char*)buffer, len);

		this->recvHandle(header, cache, this->context);
		this->dataLen = 0;
		cache.Reset();
		return 0;
	}

	enum util::endian_type endian;
	unsigned int headLen;

	void* context;
	void (*recvHandle)(DataHeader&, buffer::Data&, void*);

private:
	unsigned int dataLen;
	DataHeader header;
	buffer::Data cache;
};

#endif