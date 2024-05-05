import React, {FC} from 'react';
import './styles.scss';

interface PageTitleProps {
    title: React.ReactNode;
}

export const PageTitle: FC<PageTitleProps> = ({title}) => {
    return (
        <h1 className="page-title">
            {title}
        </h1>
    );
};