/** 
 * @file exdata.h
 * @brief 数据定义
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 描述接口数据  
 */
#ifndef ASSIST_DATA_EXDATA_H
#define ASSIST_DATA_EXDATA_H

#include <string>
#include <vector>

#include "types.h"
#include "err.h"

#include "../util/buffer.h"

// 16 + (n+2)
// HeaderVersion(4) + length(4)(n+2) + reqid(4) + cmdid(4) + reqdata(n) + checksum(2) 
const unsigned int HeaderVersion = 0x12345678;

struct DataHeader {
    unsigned int length;
    unsigned int reqid;
    unsigned int cmdid;
};

class IRequestResponse {
    public:
        virtual unsigned int length() = 0;
        virtual int serialize(buffer::Data& stream) = 0;
        virtual int deserialize(buffer::Data& stream) = 0;
};

int WriteRequest(const DataHeader& header, IRequestResponse* req, buffer::Data& stream)
{
    // HeaderVersion(4)
    stream << HeaderVersion;

    // length(4)
    stream << header.length + 2;
    // reqid(4)
    stream << header.reqid;
    // cmdid(4)
    stream << header.cmdid;

    // reqdata
    int ret = req->serialize(stream);
    if (ret != 0){
        return ret;
    }

    // checksum(2) 
    stream << "22";

    return 0;
}

int ReadResponseHeader(buffer::Data& stream, DataHeader& header)
{
    // HeaderVersion(4)
    unsigned int ver;
    stream >> ver;
    if (ver != HeaderVersion){
        return ECI_Header_Version;
    }

    // length(4)
    stream >> header.length;
    // reqid(4)
    stream >> header.reqid;
    // cmdid(4)
    stream >> header.cmdid;
    return 0;
}

int ReadResponseData(buffer::Data& stream, DataHeader& header, IRequestResponse* res)
{
     // resdata
    int ret = res->deserialize(stream);
    if (ret != 0){
        return ret;
    }
    return 0;
}

/// 心跳
struct Req_KeepAlive{
	unsigned int content;

    virtual unsigned int length(){
        return 4;
    }
    virtual int serialize(buffer::Data& stream){
        stream.writeVal(content);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
         int ret = 0;
        ret = stream.readVal(content);
        if (ret != 0){
            return ret;
        }

        return 0;
    }
};
struct Ack_KeepAlive{
	unsigned int content;

    virtual unsigned int length(){
        return 4;
    }
    virtual int serialize(buffer::Data& stream){
        stream.writeVal(content);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
         int ret = 0;
        ret = stream.readVal(content);
        if (ret != 0){
            return ret;
        }

        return 0;
    }
};

/// 公用数据
// 分组信息
struct GroupInfo : public IRequestResponse {
	std::string name;       //32bytes
    std::string comment;    //64bytes

    virtual unsigned int length(){
        return 32 + 64;
    }
    virtual int serialize(buffer::Data& stream){
        int nl = this->name.size();
        if (nl == 0) {
            return ECI_Name_Empty;
        }
        if (nl > 32) {
            return ECI_NameLength_ToLong32;
        }

        int cl = this->comment.size();
        if (cl > 64) {
            return ECI_CommentLength_ToLong64;
        }

        stream.writeVal(name, 32);
        stream.writeVal(comment, 64);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        int ret = 0;
        ret = stream.readVal(name, 32);
        if (ret != 0){
            return ret;
        }
        ret = stream.readVal(comment, 64);
        if (ret != 0){
            return ret;
        }
        return 0;
    }
};
// 应答信息
struct BaseResponse : public IRequestResponse {
	unsigned short code;
    std::string message;    //32bytes

    virtual unsigned int length(){
        return 2 + 32;
    }
    virtual int serialize(buffer::Data& stream){
        int ml = this->message.size();
        if (ml > 32) {
            return ECI_MessageLength_ToLong32;
        }

        stream.writeVal(code);
        stream.writeVal(message, 32);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        int ret = 0;
        ret = stream.readVal(code);
        if (ret != 0){
            return ret;
        }

        ret = stream.readVal(message, 32);
        if (ret != 0){
            return ret;
        }
        return 0;
    }
};
///

/// 添加分组
struct Req_GroupAdd : public IRequestResponse {
	GroupInfo groupInfo;

    virtual unsigned int length(){
        return groupInfo.length();
    }
    virtual int serialize(buffer::Data& stream){
        return groupInfo.serialize(stream);
    }
    virtual int deserialize(buffer::Data& stream){
        return groupInfo.deserialize(stream);
    }
};
struct Ack_GroupAdd : public IRequestResponse {
	BaseResponse baseResponse;

    unsigned char groupid;

    virtual unsigned int length(){
        return baseResponse.length() + 1;
    }
    virtual int serialize(buffer::Data& stream){
        int ret = baseResponse.serialize(stream);
        if (ret != 0){
            return ret;
        }

        stream.writeVal(groupid);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        int ret = baseResponse.deserialize(stream);
        if (ret != 0){
            return ret;
        }

        ret = stream.readVal(groupid);
        if (ret != 0){
            return ret;
        }
        return 0;
    }
};

/// 修改分组
struct Req_GroupMod : public IRequestResponse{
	unsigned char groupid;

    Req_GroupAdd groupAdd;
    virtual unsigned int length(){
        return 1 + groupAdd.length();
    }
    virtual int serialize(buffer::Data& stream){
        stream.writeVal(groupid);
        return groupAdd.serialize(stream);
    }
    virtual int deserialize(buffer::Data& stream){
        stream.readVal(groupid);
        return groupAdd.deserialize(stream);
    }
};
struct Ack_GroupMod{
	BaseResponse baseResponse;

    virtual unsigned int length(){
        return baseResponse.length();
    }
    virtual int serialize(buffer::Data& stream){
        int ret = baseResponse.serialize(stream);
        if (ret != 0){
            return ret;
        }

        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        int ret = baseResponse.deserialize(stream);
        if (ret != 0){
            return ret;
        }
        return 0;
    }
};

/// 删除分组
struct Req_GroupDel{
	unsigned char groupid;

    virtual unsigned int length(){
        return 1;
    }
    virtual int serialize(buffer::Data& stream){
        stream.writeVal(groupid);
        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        return stream.readVal(groupid);
    }
};
struct Ack_GroupDel{
	BaseResponse baseResponse;

    virtual unsigned int length(){
        return baseResponse.length();
    }
    virtual int serialize(buffer::Data& stream){
        int ret = baseResponse.serialize(stream);
        if (ret != 0){
            return ret;
        }

        return 0;
    }
    virtual int deserialize(buffer::Data& stream){
        int ret = baseResponse.deserialize(stream);
        if (ret != 0){
            return ret;
        }
        return 0;
    }
};

#endif