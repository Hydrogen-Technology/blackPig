import { Form,Select,Button,Flex,TimePicker,DatePicker,Card, Col, Row  } from "antd";
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import dayjs,{ Dayjs } from "dayjs";
import { useState } from "react";


const weekSetting = [
    {value:1,label:'星期一'},
    {value:2,label:'星期二'},
    {value:3,label:'星期三'},
    {value:4,label:'星期四'},
    {value:5,label:'星期五'},
    {value:6,label:'星期六'},
    {value:7,label:'星期日'},
]
interface DayType {
    day: Dayjs, 
    timeList: Array<{start:Dayjs,end:Dayjs}>
}
const More: React.FC<{defaultStartTime: Dayjs; defaultEndTime: Dayjs}> = ({defaultStartTime, defaultEndTime}) => {
    const [timeList,setTimeList] = useState<Array<{start:Dayjs,end:Dayjs}>>([])
    const [dayList, setDayList] = useState<Array<DayType>>([])
    return (
        <>
        <h1>临时设置</h1>
        <Form
            name="more"
            wrapperCol={{ span: 20 }}
            style={{ maxWidth: 600 }}
        >
            <Flex gap={'large'} style={{marginBottom:20}}>
                <Button type="default" icon={<PlusOutlined />} 
                    onClick={()=>{setDayList([...dayList,{day:dayjs(new Date()),timeList:[]}])}}
                    >添加临时日期</Button>
                <Button type="default" icon={<MinusCircleOutlined />}
                    onClick={()=>{setDayList(dayList.slice(0,dayList.length - 1))}}
                    >移除临时日期</Button>
            </Flex>
            {
                dayList.length >0? (
                    <Flex justify={"space-between"} gap={'large'} wrap>
                    {
                        dayList.map((item,index) => (
                            <Card key={index} title={item.day.format('YYYY年MM月DD日')}>
                                <Flex gap={'large'}>
                                    <Form.Item label="日期选择" name={item.day.toString()} rules={[{ required: true, message: '请选择日期!' }]} >
                                        <DatePicker onChange={(date) => {
                                            if(date) {
                                                item.day = date
                                            }
                                        }} defaultValue={item.day} />
                                    </Form.Item>
                                    <Button type="default" icon={<PlusOutlined />} 
                                        onClick={()=>{
                                            item.timeList = [...item.timeList,{start:defaultStartTime,end:defaultEndTime}]
                                            const newDayList = [...dayList]
                                            setDayList(newDayList)
                                        }
                                    }
                                        >添加时间段</Button>
                                    <Button type="default" icon={<MinusCircleOutlined />}
                                        onClick={()=>{
                                            item.timeList = item.timeList.slice(0,item.timeList.length - 1)
                                            const newDayList = [...dayList]
                                            setDayList(newDayList)
                                        }}
                                        >移除时间段</Button>
                                </Flex>
                                {
                                    item.timeList.length >0 ? (
                                        <Form.Item
                                            label="预约时间段"
                                            name="time"
                                            rules={[{ required: true, message: '请输入预约时间段!' }]}
                                        >
                                        
                                        {
                                            item.timeList.map((ii,index) => (
                                                <TimePicker.RangePicker key={index} defaultValue={[ii.start,ii.end]} style={{marginBottom:20}} />
                                            ))
                                        }
                                        </Form.Item>
                                    ):(
                                        <></>    
                                    )
                                }
                            </Card>
                            
                        ))
                    }
                    </Flex>
                ):(
                    <></>
                )
            }


            <Form.Item
                name="submit"
            >
                <Button type="primary" htmlType="submit" style={{marginTop:20}}>提交保存</Button>
            </Form.Item>  
        </Form>
        </>
    )
}

export default More;