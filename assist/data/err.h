/** 
 * @file err.h
 * @brief 错误码定义
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 描述接口错误码
 */
#ifndef ASSIST_DATA_ERR_H
#define ASSIST_DATA_ERR_H

/// 接口错误码定义
enum Err_Code{
	EC_Ok							= 0		,///< 0：成功
};

/// 内部错误码定义
enum Err_Code_Internal{
	ECI_Ok							= 0		,///< 0：成功
	ECI_Header_Version				= 1		,///< 1：头版本号错误
	ECI_Name_Empty					= 2		,///< 2：名称为空
	ECI_NameLength_ToLong32			= 3		,///< 3：名称长度超过32个字符
	ECI_CommentLength_ToLong64		= 4		,///< 4：备注长度超过64个字符
	ECI_MessageLength_ToLong32		= 5		,///< 5：消息长度超过32个字符
};

static const char* Err_Code_Internal_Msg[] = {
	"成功",
	"头版本号错误",
	"名称为空",
	"名称长度超过32个字符",
	"备注长度超过64个字符",
	"消息长度超过32个字符",
};

static const char* Get_Err_Code_Internal_Msg(int code){
	if (code < sizeof(Err_Code_Internal_Msg)/sizeof(const char*)){
		return Err_Code_Internal_Msg[code];
	}

	return "未知错误";
}

#endif