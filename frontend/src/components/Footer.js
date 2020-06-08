import React from 'react';
import {Layout} from 'antd';

const {Footer} = Layout;

const MyFooter = () => {
    return (
        <Footer style={{
            textAlign: 'right',
            position: 'fixed',
            width: '100%',
            bottom: 0,
            backgroundColor: '#112d4e',
            color: '#dbe2ef',
            fontWeight: 500
        }}>
            <a href={'https://github.com/adamdevigili/skillbased.io'}>Created by Adam Devigili</a>
        </Footer>
    )
}


export default MyFooter;
