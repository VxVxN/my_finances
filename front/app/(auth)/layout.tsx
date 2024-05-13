import { FC } from 'react'
import './styles.scss'

interface LayoutDashboardProps {
	children: React.ReactNode
}

const LayoutDashboard: FC<LayoutDashboardProps> = ({ children }) => {
	return (
		<div className='auth layout'>
			<main>{children}</main>
		</div>
	)
}

export default LayoutDashboard
