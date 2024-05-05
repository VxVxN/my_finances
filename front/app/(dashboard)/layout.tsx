import {FC} from "react";
import {NavBar} from '@/widgets/navbar';

interface LayoutDashboardProps {
    children: React.ReactNode;
}

const LayoutDashboard: FC<LayoutDashboardProps> = ({children}) => {
    return <div>
        <NavBar/>
        <main>
            {children}
        </main>
        <footer>

        </footer>
    </div>;
};

export default LayoutDashboard;
