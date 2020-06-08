import React, {useState} from 'react';
import SportsDropdown from './components/Dropdown';
import GoButton from './components/GoButton';
import HomepageLayout from "./components/HomepageLayout";
import MyTitle from "./components/Title";
import { Layout, Menu, Breadcrumb } from 'antd';
import './App.less';
import MyFooter from "./components/Footer";

const { Header, Footer, Content } = Layout;

function App() {
    return (
        <Layout className="layout">
            <Header style={{
                position: 'fixed',
                zIndex: 1,
                width: '100%'
            }}>
                <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['1']}>
                    <Menu.Item key="1">Home</Menu.Item>
                    <Menu.Item key="2">Sports</Menu.Item>
                    <Menu.Item key="3">Teams</Menu.Item>
                    <Menu.Item key="4">Players</Menu.Item>
                    <Menu.Item key="5">Account</Menu.Item>
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
                    <MyTitle title="SKILLBASED.IO"/>
                    <h3>Select a sport to get started</h3>
                    <div className="sport_dropdown">
                        <SportsDropdown />
                        <div className="go_button">
                            <GoButton />
                        </div>
                    </div>
                </div>
            </Content>
            <MyFooter/>
        </Layout>
    );
}

export default App;
