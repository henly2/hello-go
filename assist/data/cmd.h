/** 
 * @file cmd.h
 * @brief 指令定义
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 描述接口指令  
 */
#ifndef ASSIST_DATA_CMD_H
#define ASSIST_DATA_CMD_H

/// 指令定义
enum Command{
	Command_KeepAlive	        =  0x0000000E, ///< 心跳

    Command_GroupAdd	        =  0x00000101, ///< 增加分组
    Command_GroupMod	        =  0x00000102, ///< 修改分组
    Command_GroupDel	        =  0x00000103, ///< 删除分组
    Command_GroupRead	        =  0x00000104, ///< 读取分组

    Command_UserReg	            =  0x00000105, ///< 人员注册
    Command_UserMod	            =  0x00000106, ///< 人员修改
    Command_UserDel	            =  0x00000107, ///< 人员删除
    Command_UserReadByGroup	    =  0x00000108, ///< 人员读取（按组读取）

    Command_UserCodeAdd	        =  0x00000109, ///< 人员指纹增加
    Command_UserCodeMod	        =  0x0000010A, ///< 人员指纹修改
    Command_UserCodeDel	        =  0x0000010B, ///< 人员指纹删除
    Command_UserCodeRead	    =  0x0000010C, ///< 人员指纹查询

    Command_UserPhotoAdd	    =  0x0000010D, ///< 人员照片增加
    Command_UserPhotoMod	    =  0x0000010E, ///< 人员照片修改
    Command_UserPhotoDel	    =  0x0000010F, ///< 人员照片删除
    Command_UserPhotoRead	    =  0x00000110, ///< 人员照片查询

    Command_DevReg              =  0x00000111, ///< 设备注册
    Command_DevMod              =  0x00000112, ///< 设备修改
    Command_DevDel              =  0x00000113, ///< 设备删除
    Command_DevRead             =  0x00000114, ///< 设备读取
};

#endif