'use client'

import { Input } from '@nextui-org/react'
import { useFormik } from 'formik'
import * as Yup from 'yup'

type SignUpForm = {
	login: string
	password: string
	repeatPassword: string
}

const validationSchema = Yup.object().shape({
	login: Yup.string().required('Логин обязателен'),
	password: Yup.string().required('Пароль обязателен'),
	repeatPassword: Yup.string().oneOf(
		[Yup.ref('password')],
		'Пароли должны совпадать'
	),
})

export const SignUpForm = () => {
	const formik = useFormik<SignUpForm>({
		initialValues: {
			login: '',
			password: '',
			repeatPassword: '',
		},
		validationSchema,
		onSubmit: values => {
			alert(JSON.stringify(values, null, 2))
		},
	})

	return (
		<div>
			<form onSubmit={formik.handleSubmit} className='flex flex-col gap-4'>
				<Input
					type='text'
					placeholder='Логин'
					isInvalid={!!formik.errors.login}
					errorMessage={formik.errors.login}
				/>
				<Input
					type='password'
					placeholder='Пароль'
					errorMessage={formik.errors.password}
				/>
				<Input
					type='password'
					placeholder='Повторите пароль'
					errorMessage={formik.errors.repeatPassword}
				/>
			</form>
		</div>
	)
}
