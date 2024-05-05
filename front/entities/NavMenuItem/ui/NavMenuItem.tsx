'use client';

import React, {FC} from 'react';
import Link from 'next/link';
import {NavMenuType} from '@/entities/NavMenuItem';
import {usePathname} from 'next/navigation';
import './styles.scss'

interface NavMenuItemProps {
    data: NavMenuType;
}

export const NavMenuItem: FC<NavMenuItemProps> = ({data}) => {
    const pathname = usePathname()

    return (
        <Link href={data.link} className={`nav-link-item ${pathname === data.link ? 'active' : ''}`}>
            {data.title}
        </Link>
    );
};
