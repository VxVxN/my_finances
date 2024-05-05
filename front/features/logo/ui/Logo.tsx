import React, {FC} from 'react';
import Link from 'next/link';
import {LogoIcon} from '@/shared/icons';
import './styles.scss'



export const Logo: FC = () => {
    return (
        <Link className='logo' href={"/"}>
            <LogoIcon/><div className='title'>My Finances</div>
        </Link>
    );
};