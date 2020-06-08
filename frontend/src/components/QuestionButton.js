import React from 'react';
import { Button } from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';


const QuestionButton = () => {
    return (
        <Button
            icon={ <QuestionCircleOutlined />}
            shape="circle"
            ghost="true"
            size='large'
        />
    )
}


export default QuestionButton;
