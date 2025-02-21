import React from 'react'
import { View,Image } from '@tarojs/components'

import './index.scss'
import { Container } from "../../components/container/container"
// @ts-ignore
import blue_cat from "../../assets/images/蓝猫.jpg"
// @ts-ignore
import red_cat from "../../assets/images/红猫.jpg"

export default function Index() {
  return (
    <>
    <Container>
      <View className='img' ></View>
      <View className='top'>
        <View className='title'>约动场</View>
      </View>
      <View className='main'>
        <View className='card'  onClick={getReserve}>
          <View className='text'>我是领导</View>
          <Image className='icon' src={red_cat} />
        </View>
        <View className='card' onClick={setReserve}>
          <View className='text' >我要预约</View>
          <Image className='icon' src={blue_cat} />
        </View>
      </View>
    </Container>
    </>
  )
}
function getReserve() {
  console.log('getReserve,跳转')
}
function setReserve() {
  console.log('setReserve')
}