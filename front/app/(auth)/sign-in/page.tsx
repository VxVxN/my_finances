 
import SignInForm from '@/widgets/SignInForm/ui/SignInForm'
import '../styles.scss'

export default function SignInPage() {
	return (
		<div className="wrapper-auth">
			<h1>Авторизация</h1>
			<SignInForm />
		</div>
	)
}
