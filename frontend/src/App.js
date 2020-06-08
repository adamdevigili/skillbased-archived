import React, {useState} from 'react';
import SportsDropdown from './components/Dropdown';
import GoButton from './components/GoButton';
import MyTitle from "./components/Title";
import { Layout, Menu, Breadcrumb } from 'antd';
import './App.less';
import MyFooter from "./components/Footer";
import { HomeOutlined, UserOutlined, QuestionCircleOutlined } from '@ant-design/icons';
import QuestionButton from "./components/QuestionButton";

const SubMenu = Menu.SubMenu;
const MenuItemGroup = Menu.ItemGroup;

const { Header, Footer, Content } = Layout;

function App() {
    return (
        <Layout className="layout">
            <Header style={{
                position: 'fixed',
                width: '100%'
            }}>
                <Menu
                    theme="dark"
                    mode="horizontal"
                    defaultSelectedKeys={['home']}
                >
                    <Menu.Item style={{float: 'left'}} key="home" icon={<HomeOutlined />}/>
                    <Menu.Item style={{float: 'left'}} key="sports">Sports</Menu.Item>
                    <Menu.Item style={{float: 'left'}} key="teams">Teams</Menu.Item>
                    <Menu.Item style={{float: 'left'}} key="players">Players</Menu.Item>

                    <Menu.Item style={{float: 'right'}} key="account" icon={<UserOutlined />}/>
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
                        <QuestionButton />
                    </div>
                </div>
            </Content>
            <MyFooter/>
        </Layout>
    );
}

export default App;
