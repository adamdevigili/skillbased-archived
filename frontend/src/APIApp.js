import React from 'react';
import './App.less';
import {Layout, Menu} from "antd";
import MyTitle from "./components/Title";
import SportsDropdown from "./components/Dropdown";
import GoButton from "./components/GoButton";
import MyFooter from "./components/Footer";

const { Header, Footer, Content } = Layout;

function APIApp() {
    return (
        <Layout className="layout">
            <Header style={{
                position: 'fixed',
                zIndex: 1,
                width: '100%'
            }}>
                <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['1']}>
                    <Menu.Item key="1">Home</Menu.Item>
                    <Menu.Item key="2">Docs</Menu.Item>
                    <Menu.Item key="3">Account</Menu.Item>
                </Menu>
            </Header>

            <Content style={{
                textAlign: 'center',
                alignItems: 'center',
                display: 'flex',
                justifyContent: 'center',
                height: '100vh',
                backgroundColor: '#112d4e'

            }}>
                <div className="site-layout-content">
                    <MyTitle title="API.SKILLBASED.IO"/>
                    <h3>In Development</h3>
                </div>
            </Content>
            <MyFooter/>
        </Layout>
    );
}

export default APIApp;
