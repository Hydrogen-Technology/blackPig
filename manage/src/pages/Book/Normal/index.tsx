import { Divider,Form,Select,Button,Flex,Switch,TimePicker  } from "antd";
import { MinusCircleOutlined, PlusOutlined,CheckOutlined, CloseOutlined } from '@ant-design/icons';
import { useEffect,useState } from "react";
import dayjs, { Dayjs } from 'dayjs';


const dayLine = [
    {value:2,label:'2天内'},
    {value:3,label:'3天内'},
    {value:4,label:'4天内'},
    {value:5,label:'5天内'},
    {value:6,label:'6天内'},
    {value:7,label:'7天内'},
    {value:15,label:'15天内'},
    {value:30,label:'30天内'},
]
const Normal: React.FC<{defaultStartTime:Dayjs,defaultEndTime:Dayjs}> = ({defaultStartTime,defaultEndTime}) => {
    const [TimeList,setTimeList] = useState([{
        start:defaultStartTime,
        end:defaultEndTime,
    }])
    
    const addTime = () => {
        setTimeList([...TimeList,{
            start:defaultStartTime,
            end:defaultEndTime, 
        }])
    };
    const removeTime = () => {
        if(TimeList.length === 1) return;
        setTimeList(TimeList.slice(0,TimeList.length - 1)) 
    }
    return (
        <>
        <h1>预约设置</h1>
        <Form
        name="book"

        wrapperCol={{ span: 20 }}
        style={{ maxWidth: 600 }}
    >
        <Flex vertical
            justify={'space-between'}
        >   
            <Flex gap={'large'}>
                <Form.Item
                    label="预约期限"
                    name="bookLine"
                    rules={[{ required: true, message: '请输入预约期限!' }]}
                >
                    <Select defaultValue={7} style={{width:150}} options={dayLine}>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="避免周末"
                    name="weekend"
                    labelCol={{ span: 15 }}
                >
                    <Switch defaultChecked 
                    checkedChildren={<CheckOutlined />}
                    unCheckedChildren={<CloseOutlined />}
                    style={{marginLeft:20}} />
                </Form.Item>
                <Form.Item
                    label="避免节假日"
                    name="holiday"
                    labelCol={{ span: 15 }}
                >
                    <Switch defaultChecked 
                    checkedChildren={<CheckOutlined />}
                    unCheckedChildren={<CloseOutlined />}
                    style={{marginLeft:20}} />
                </Form.Item>
            </Flex>
            <Flex gap={'large'} style={{marginBottom:20}}>
                <Button type="default" icon={<PlusOutlined />} onClick={addTime}>添加时间段</Button>
                <Button type="default" icon={<MinusCircleOutlined />} onClick={removeTime}>移除时间段</Button>
            </Flex>
            <Form.Item
                label="预约时间段"
                name="time"
                rules={[{ required: true, message: '请输入预约时间段!' }]}
            >  
            {
                TimeList.map((item,index) => (
                    <TimePicker.RangePicker defaultValue={[item.start,item.end]} style={{marginBottom:20}} />
                ))
            }
            </Form.Item>
            <Form.Item
                name="submit"
            >
                <Button type="primary" htmlType="submit">提交保存</Button>
            </Form.Item>
        </Flex>
    </Form>
    </>
    )
}


export default Normal