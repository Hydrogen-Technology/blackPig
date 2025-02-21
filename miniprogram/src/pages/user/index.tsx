import { AtAvatar,AtList, AtListItem,AtGrid } from "taro-ui"
import React from "react"
import { View } from "@tarojs/components"
import { Container } from "../../components/container/container"
import "./index.scss"


const Demo1 = () => {
  return (
    <>
    <Container>
      <View className='container at-row at-row__align--center'>
        <AtAvatar circle text='微信用户'></AtAvatar>
        <View className='at-col right-message'>
          <View className='at-row'>微信用户</View>
          <View className='at-row description'>微信用户，欢迎你</View>
        </View>
      </View>
      <View className='at-col button-list'>
      <AtGrid data={
          [
            {
              image: 'https://img12.360buyimg.com/jdphoto/s72x72_jfs/t6160/14/2008729947/2754/7d512a86/595c3aeeNa89ddf71.png',
              value: '我的预约'
            },
            {
              image: 'https://img20.360buyimg.com/jdphoto/s72x72_jfs/t15151/308/1012305375/2300/536ee6ef/5a411466N040a074b.png',
              value: '历史记录'
            },
            {
              image: 'https://img10.360buyimg.com/jdphoto/s72x72_jfs/t5872/209/5240187906/2872/8fa98cd/595c3b2aN4155b931.png',
              value: '预约设置'
            }
          ]
      }
      />
      </View>
      <View  className='at-col button-list' >
        <AtList hasBorder={false}>
          <AtListItem title='联系我们' arrow='right' />
          <AtListItem title='使用条款' arrow='right' />
          <AtListItem title='个人信息保护政策' arrow='right' />
        </AtList>
      </View>
      <View  className='at-col button-list' >
        <AtList>
          <AtListItem title='退出登录' arrow='right' />
        </AtList>
      </View>
    </Container>
    </>
  )
}
export default Demo1
