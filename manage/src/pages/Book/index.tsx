import { Divider,Form,Select,Button,Flex,Switch,TimePicker  } from "antd";
import { MinusCircleOutlined, PlusOutlined,CheckOutlined, CloseOutlined } from '@ant-design/icons';
import { useEffect,useState } from "react";
import dayjs, { Dayjs } from 'dayjs';
import More from "./More/idex";
import Normal from "./Normal";
import Temp from "./Temp";


const defaultStartTime = dayjs('00:00:00','HH:mm:ss')
const defaultEndTime = dayjs('23:59:59','HH:mm:ss')
const BookList : React.FC = () => {
    return (
        <div>
            <Normal defaultStartTime={defaultStartTime} defaultEndTime={defaultEndTime}></Normal>
            <Divider />
            <More defaultStartTime={defaultStartTime} defaultEndTime={defaultEndTime}></More>
            <Divider />
            <Temp defaultStartTime={defaultStartTime} defaultEndTime={defaultEndTime}></Temp>
        </div>
    )
}

export default BookList;