import React, {FC} from 'react';
import {
    Button,
    Navbar,
    NavbarBrand,
    NavbarContent,
    NavbarItem,
    NavbarMenu,
    NavbarMenuItem,
    NavbarMenuToggle
} from '@nextui-org/react';
import Link from 'next/link';
import {Logo} from '@/features/logo';
import {NavMenuItem, NavMenuType} from '@/entities/NavMenuItem';
import './styles.scss'
import {SwitchTheme} from '@/features/SwitchTheme';
import {LoginButton} from '@/features/LoginButton';

interface NavBarProps {

}

const menuItems: NavMenuType[] = [
    {
        title: 'Сделки',
        link: '/orders'
    },
    {
        title: 'Статистика',
        link: '/statistics'
    },

];

export const NavBar: FC<NavBarProps> = ({}) => {
    return (
        <Navbar className='nav-bar'>
            <NavbarContent className="" justify="start">
                <NavbarMenuToggle className='toggle-btn-menu'/>
                <NavbarBrand>
                    <Logo />
                </NavbarBrand>
            </NavbarContent>

            <NavbarContent className="nav-links" justify="center">
                {
                    menuItems.map((item, index) => (
                        <NavbarItem key={index}>
                            <NavMenuItem key={index}  data={item}/>
                        </NavbarItem>
                    ))
                }
            </NavbarContent>

            <NavbarContent justify="end">
                <NavbarItem>
                    <SwitchTheme/>
                </NavbarItem>
                {/*<NavbarItem className="  lg:flex">*/}
                {/*    <Link href="#">Войти</Link>*/}
                {/*</NavbarItem>*/}
                <NavbarItem>
                    <LoginButton/>
                </NavbarItem>
            </NavbarContent>

            <NavbarMenu>
                {menuItems.map((item, index) => (
                    <NavbarMenuItem key={`${item}-${index}`}>
                        <Link
                            className="w-full"
                            color={
                                index === 2 ? "warning" : index === menuItems.length - 1 ? "danger" : "foreground"
                            }
                            href="#"
                        >
                            {item.title}
                        </Link>
                    </NavbarMenuItem>
                ))}
            </NavbarMenu>
        </Navbar>
    );
};
