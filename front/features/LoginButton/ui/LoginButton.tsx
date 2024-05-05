import React, {FC} from 'react';
import {Button} from '@nextui-org/react';
import {FaKey} from 'react-icons/fa';
import './styles.scss';


export const LoginButton: FC = () => {
    return (
        <Button className="login-btn" startContent={<FaKey/>} color="primary" href="#">
            <div className="title">
                Войти
            </div>
        </Button>
    );
};