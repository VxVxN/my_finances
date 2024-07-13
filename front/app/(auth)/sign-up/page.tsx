import { SignUpForm } from '@/widgets/SignUpForm'
import '../styles.scss'

export default function SignUpPage() {
	return (
		<div className="wrapper-auth">
			<h1>Регистрация</h1>
			<SignUpForm />
		</div>
	)
}
