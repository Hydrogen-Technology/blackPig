import React, { useState, useEffect } from 'react';
import Taro from '@tarojs/taro';
import { View,Picker  } from '@tarojs/components';
import { AtCalendar,AtDivider,AtList, AtListItem,AtButton,AtTag   } from "taro-ui"
import './index.scss'
import { Container } from '../../../components/container/container';

// 已预约的时间接口规范
interface ReservedRepo {
  start_time: string,
  end_time: string,
}
export default function Index() {
  const [choosedDate, setChoosedDate] = useState(formatDate(new Date()))
  const [validDates, setValidDates] = useState<{ value: string }[]>([])
  const [timeTables,setTimeTables] = useState([<TimeTable key={0} index={0} />])
  const [reservedTimes,setReservedTimes] = useState<ReservedRepo[]>([ 
    {start_time:'10:00',end_time:'11:00'},
    {start_time:'13:00',end_time:'15:00'}
  ]) // 导员被预约
  const [submitTimes ,setSubmitTimes] = useState<{start_time:string,end_time:string}[]>([{start_time:'11:00',end_time:'10:00'}]) // 导员提交预约的时间段
  function TimeTable({index}) {
    // ? 为什么这里的index 是 {index:0}
    console.log("index:",submitTimes)
    return (
      <View className='time'>
        <Picker mode='time' value={submitTimes[index].start_time} onChange={(e)=>startTimeOnchange(e,index)} >
          <AtList>
            <AtListItem title='开始时间' extraText={submitTimes[index].start_time} />
          </AtList>
        </Picker>
        <Picker mode='time' value={submitTimes[index].end_time} onChange={(e)=>endTimeOnchange(e,index)} >
          <AtList>
            <AtListItem title='结束时间' extraText={submitTimes[index].end_time} className='end' />
          </AtList>
        </Picker>
      </View>
      
    )
  }
  function startTimeOnchange(e,index) {
    const newSubmitTimes = submitTimes.map(item => ({ ...item })) // 深拷贝对象
    console.log("start:",newSubmitTimes,newSubmitTimes[index],index)
    submitTimes[index].start_time = e.detail.value
    setSubmitTimes(newSubmitTimes)
    console.log("start:",submitTimes)
  }
  function endTimeOnchange(e,index) {
    const newSubmitTimes = [...submitTimes]
    newSubmitTimes[index].end_time = e.detail.value
    setSubmitTimes(newSubmitTimes)
    console.log("start:",submitTimes)
  }
  
  function add_time_list() {// ?在页面上新增新的timeTable
    if(timeTables.length >= 2){
      Taro.showToast({
        title: '最多只能添加2个时间段',
        icon: 'none',
        duration: 2000
      })
    }else{
      setSubmitTimes([...submitTimes, { start_time: '11:00', end_time: '10:00' }])
      setTimeTables([...timeTables,<TimeTable key={timeTables.length} index={timeTables.length} />])
    }
  }
  function dayClick(item:{value:string}) { // ?点击日期后，获取导员已经预约成功的时间段
    // ?点击日期后，获取导员已经预约成功的时间段
    setReservedTimes([
      {start_time:'9:00',end_time:'10:00'},
      {start_time:'14:00',end_time:'15:00'}])
  }
  useEffect(() => {}, [choosedDate])
  return (
    <Container>
      <AtCalendar
        className='panel'
        validDates={validDates}
        minDate={formatDate(new Date())}
        maxDate={formatDate(new Date(Date.now() + 7 * 24 * 3600 * 1000))}
        onDayClick={(e)=>dayClick(e)}
      />
      <View className='panel'>
      <AtTag className='tag'>预约情况</AtTag>
      {reservedTimes.map((item:ReservedRepo,index)=>{
        return (
          <>
            <View key={index} className='at-row at-row__justify--around'>
              <View className='at-col reserved'>开始时间：{item.start_time}</View>
              <View className='at-col reserved'>结束时间：{item.end_time}</View>
            </View>
            <AtDivider lineColor='#d01257' className='divider'>
            </AtDivider>
          </>
        )
      })}
      </View>
      <View className='bottom panel'>
        <AtTag className='tag'>预约时间段</AtTag>
        <View id='timeTable'>
          {timeTables.map((item) => {return item})}
        </View>
        <View className='at-row at-row__justify--between'>
          <AtButton type='primary' className='btn at-col at-col-5' onClick={()=>add_time_list()}>添加时间</AtButton>
          <AtButton type='secondary' className='btn at-col at-col-5'>申请预约</AtButton>
        </View>
      </View>
    </Container>
  )
}

function formatDate(date: Date): string {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  return `${year}/${month.toString().padStart(2, '0')}/${day.toString().padStart(2, '0')}`
}