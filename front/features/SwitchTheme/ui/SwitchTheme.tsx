'use client';

import React, {FC} from 'react';
import {Switch} from '@nextui-org/react';
import {MoonIcon, SunIcon} from '@nextui-org/shared-icons';
import {useTheme} from 'next-themes';

interface SwitchThemeProps {

}

export const SwitchTheme: FC<SwitchThemeProps> = ({}) => {
    const { theme, setTheme } = useTheme()

    const toggleTheme = () => {
        setTheme(theme === 'dark' ? 'light' : 'dark');
    }

    return (
        <Switch
            size="sm"
            color="primary"
            defaultSelected={theme === 'light'}
            onValueChange={toggleTheme}
            thumbIcon={({ isSelected, className }) =>
                isSelected ? (
                    <SunIcon className={className} />
                ) : (
                    <MoonIcon className={className} />
                )
            }
        >
            Dark mode
        </Switch>
    );
};