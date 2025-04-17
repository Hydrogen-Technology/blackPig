import { Form,Select,Button,Flex,TimePicker  } from "antd";
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Dayjs } from "dayjs";
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
const More: React.FC<{defaultStartTime: Dayjs; defaultEndTime: Dayjs}> = ({defaultStartTime, defaultEndTime}) => {
    const [weekTimeList,setWeekTimeList] = useState<Array<{start:Dayjs,end:Dayjs}>>([])
    return (
        <>
        <h1>高级设置</h1>
        <Form
            name="more"
            wrapperCol={{ span: 20 }}
            style={{ maxWidth: 600 }}
        >
            <Flex>
                <Form.Item
                    label="星期设置"
                    name="week"
                >
                    <Select defaultValue={[1,2,3,4,5]} style={{width:150}} options={weekSetting}>
                    </Select>
                </Form.Item>
            </Flex>
            <Flex gap={'large'} style={{marginBottom:20}}>
                <Button type="default" icon={<PlusOutlined />} 
                    onClick={()=>{setWeekTimeList([...weekTimeList,{
                        start:defaultStartTime,
                        end:defaultEndTime, 
                    }])}}
                    >添加时间段</Button>
                <Button type="default" icon={<MinusCircleOutlined />}
                    onClick={()=>{setWeekTimeList(weekTimeList.slice(0,weekTimeList.length - 1))}}
                    >移除时间段</Button>
            </Flex>
            {
                weekTimeList.length >0 ? (
                    <Form.Item
                        label="预约时间段"
                        name="time"
                        rules={[{ required: true, message: '请输入预约时间段!' }]}
                    >
                    {
                        weekTimeList.map((item,index) => (
                            <TimePicker.RangePicker defaultValue={[item.start,item.end]} style={{marginBottom:20}} />
                        ))
                    }
                    </Form.Item>
                ):(
                    <></>    
                )
            }
            <Form.Item
                name="submit"
            >
                <Button type="primary" htmlType="submit">提交保存</Button>
            </Form.Item>  
        </Form>
        </>
    )
}

export default More;