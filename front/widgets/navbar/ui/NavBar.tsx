'use client'

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
import {MdCreate} from 'react-icons/md';
import {CreateOrderForm} from '@/entities/CreateOrderForm';
import {useDisclosure} from '@nextui-org/modal';

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
    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    return (
        <>
            <Navbar className='nav-bar'>
                <NavbarContent  >
                    <NavbarMenuToggle className='toggle-btn-menu'/>
                    <NavbarBrand>
                        <Logo />
                    </NavbarBrand>
                </NavbarContent>

                <NavbarContent className="nav-links" justify="start">
                    {
                        menuItems.map((item, index) => (
                            <NavbarItem key={index}>
                                <NavMenuItem key={index}  data={item}/>
                            </NavbarItem>
                        ))
                    }
                </NavbarContent>

                <NavbarContent justify="end">
                    <NavbarItem className="  lg:flex">
                        <Button startContent={<MdCreate />} size='sm' onClick={onOpen}>Создать</Button>
                    </NavbarItem>
                    <NavbarItem>
                        <SwitchTheme/>
                    </NavbarItem>
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
            <CreateOrderForm isOpen={isOpen} onOpenChange={onOpenChange}/>
        </>
    );
};
