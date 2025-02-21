import React, { useEffect, useState } from 'react'
import { AtCard } from 'taro-ui'
import { View } from '@tarojs/components'

import './index.scss'
import { Container } from '../../components/container/container'



// !设置菜单列表
const menuList = [
  {
    id:1,
    name:'我的预约',
    path:'/reserve/index'
  },
  {
    id:2,
    name:'进行中',
    path:'/reserve/pending/index'
  },
  {
    id:3,
    name:'已完成',
    path:'/reserve/completed/index'
  },
  {
    id:4,
    name:'已取消',
    path:'/reserve/canceled/index'
  }
]
export default function Index () {
  const [reserveList , setReserveList] = useState<reserveItem[]>([])
  const [currentSelected,setCurrentSelected] = useState(1)
  useEffect(()=>{
    getServeList({reserveList,setReserveList})
  },[])
  return (
    <Container>
      <View className='menu-panel'>
        <Menu currentSelected={currentSelected} setCurrentSelected={setCurrentSelected}  />
      </View>
      <View className='content'>
        <HistoryList reserveList={reserveList} />
      </View>
    </Container>
  )
  
}
const menuSelected = (id,setCurrentSelected) => {
  console.log(id)
  setCurrentSelected(id)
  changeReserveList(id)
}
const Menu = ({currentSelected,setCurrentSelected}) => {
  return (
    <>
      {menuList.map((item) => (
        <View key={item.id} onClick={menuSelected.bind(null,item.id,setCurrentSelected)} className={`${item.id === currentSelected ? 'selected':'' }`} >{item.name}</View>
      ))}
    </>
  )
  
}
const getServeList = ({reserveList,setReserveList})=> {
  setReserveList([
    {
      id: '1',
      name: '张三',
      avatar: '',
      created_time: '2024-12-12 12:12:12',
      reserve_time: '2024-12-19 12-13',
      state: '进行中'
    },
    {
      id: '2',
      name: '李斯',
      avatar: '',
      created_time: '2024-12-12 12:12:12',
      reserve_time: '2024-12-19 12-13',
      state: '已完成'
    },
    {
      id: '3',
      name: '马素超',
      avatar: '',
      created_time: '2024-12-15 12:10:12',
      reserve_time: '2024-12-19 12-13',
      state: '已取消'
    },
  ])
  console.log(reserveList)
}
interface changeReserveListProps {
  id:number,
  reserveList:reserveItem[],
  setReserveList:(list :reserveItem[]) => void
}
const changeReserveList = ({id,reserveList,setReserveList}:changeReserveListProps) => {
  const temps:reserveItem[] = []
  const states = ['所有预约','进行中','已完成','已取消']
  // ! 这里暂时写死，每个id代表对应的状态
  console.log(reserveList)
  reserveList.map((item:reserveItem)=> (
     states[id-1] === item.state ? temps.push(item) : null
  ))
  setReserveList(temps)
}

interface HistoryListProps {
  reserveList:reserveItem[]
}
const HistoryList = ({reserveList}:HistoryListProps) => {
  // todo 要请求从后端获取列表any信息，这里需要传递参数，表示是已预约或者其他
  return (
    <>
    {
      reserveList.map((item)=>(
          <AtCard key={item.id} className='reserve-card'
            title={item.name}
            note={`下单时间：${item.created_time}`}
            extra={item.state}
            thumb='https://s1.aigei.com/src/img/png/14/141ef919be0943a88701c0693471b09e.png?imageMogr2/auto-orient/thumbnail/!282x282r/gravity/Center/crop/282x282/quality/85/%7CimageView2/2/w/282&e=2051020800&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:vMzBohrKoYn3Ss14rhzj0OxuYzg='
          >
            预约时间:{item.reserve_time}
          </AtCard>
      ))
    }
    </>
   )   
}
// 预约卡片接口定义
interface reserveItem {
  id:string,
  name:string,
  avatar:string,
  created_time:string,
  reserve_time:string,
  state:string
}




