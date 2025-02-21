import React,{useState,useEffect} from 'react';
import { View } from '@tarojs/components';
import { AtList, AtListItem,AtSearchBar } from 'taro-ui';
import './index.scss'
import { Container } from '../../components/container/container';



const leaderList = [
  {
    id:'1',
    name: '张三导员',
    phone: '13888888888',
    avatar: 'http://img12.360buyimg.com/jdphoto/s72x72_jfs/t10660/330/203667368/1672/801735d7/59c85643N31e68303.png'
  },
  {
    id:'2',
    name: '李四导员',
    phone: '13888888888',
    avatar: 'http://img12.360buyimg.com/jdphoto/s72x72_jfs/t10660/330/203667368/1672/801735d7/59c85643N31e68303.png'
  },
  {
    id:'3',
    name: '王五导员',
    phone: '13888888888',
    avatar: 'http://img12.360buyimg.com/jdphoto/s72x72_jfs/t10660/330/203667368/1672/801735d7/59c85643N31e68303.png'
  },
]
/**
 * Index 组件用于渲染预约页面。
 * 
 * @returns {JSX.Element} 返回包含预约文本的视图组件。
 */
export default function Index() {
  const [searchValue,setSearchValue] = useState('')
  const searchLeaders = (value) => {
    console.log("搜索内容：",value)
    setSearchValue(value)
  }
  return (
    <>
      <Container>
        <AtSearchBar
          value={searchValue}
          onChange={searchLeaders}
          actionName='搜索'
          showActionButton
        />
        <AtList className='base'>
          <Leaders />
        </AtList>
      </Container>
    </>
  )
}

const Leaders = () => {
  return(
    <>
    {leaderList.map((item)=>(
      <AtListItem className='leaders'
        key={item.id}
        title={item.name}
        thumb={item.avatar}
        onClick={reserve.bind(null,item.id)}
        arrow='right'
        note={`联系电话： ${item.phone}`}
      />
    ))
    }
    </>
  )
}
const reserve = (id:string) => {
  console.log('预约函数跳转',id)
}