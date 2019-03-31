/** 
 * @file endian.h
 * @brief 大小端 
 * @sa null
 * @author liuheng
 * @date 3/30/2019
 *
 * 大小端 
 */
#ifndef ASSIST_UTIL_ENDIAN_H
#define ASSIST_UTIL_ENDIAN_H

#include <stdio.h>

namespace util
{

enum endian_type
{
	EXBUFFER_BIG_ENDIAN,
	EXBUFFER_LITTLE_ENDIAN
};

/*
本地 s_host_endian 
> 最后一位表示初始化位
> 前一位表示类型位，0:big_endian; 1:little_endian
*/
static unsigned char s_host_endian = 0;
static void init_host_endian()
{
	if((s_host_endian & 0x01) == 0x01)
	{
		return;
	}

	s_host_endian = s_host_endian|0x01;
	unsigned short int i =0x1234;
	unsigned char* p = (unsigned char*)&i;
	if(*p == 0x12)
	{
		s_host_endian = s_host_endian ^ (~s_host_endian & (0x00<<1));
        // s_host_endian = 0x01
        printf(">>host big endian\n");
	}
	else
	{
		s_host_endian = s_host_endian ^ (~s_host_endian & (0x01<<1));
        // s_host_endian = 0x03
        printf(">>host little endian\n");
	}
}

static bool is_host_endian_little()
{
    init_host_endian();

	return (s_host_endian & (0x01<<1)) == 0x02;
}

/*
本地与网络转换
*/
static unsigned long exchange(unsigned long x, bool is_to_little)
{
	if(is_host_endian_little() == is_to_little)
	{
		return x;
	}

	return
	((unsigned long)(	\
        (((unsigned long)(x) & (unsigned long)0x000000ffUL) << 24) |	 \
        (((unsigned long)(x) & (unsigned long)0x0000ff00UL) <<  8) |	 \
        (((unsigned long)(x) & (unsigned long)0x00ff0000UL) >>  8) |	 \
        (((unsigned long)(x) & (unsigned long)0xff000000UL) >> 24)));
}

static unsigned long exchange(unsigned int x, bool is_to_little)
{
	if(is_host_endian_little() == is_to_little)
	{
		return x;
	}

	return
	((unsigned long)(	\
        (((unsigned long)(x) & (unsigned long)0x000000ffUL) << 24) |	 \
        (((unsigned long)(x) & (unsigned long)0x0000ff00UL) <<  8) |	 \
        (((unsigned long)(x) & (unsigned long)0x00ff0000UL) >>  8) |	 \
        (((unsigned long)(x) & (unsigned long)0xff000000UL) >> 24)));
}

static unsigned short exchange(unsigned short x, bool is_to_little)
{
	if(is_host_endian_little() == is_to_little)
	{
		return x;
	}

	return
	((unsigned short)(	\
        (((unsigned short)(x) & (unsigned short)0x00ffU) << 8) |	 \
        (((unsigned short)(x) & (unsigned short)0xff00U) >> 8)));
}

} // end of namespace util

#endif