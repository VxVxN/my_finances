'use client'

import { signUp } from '@/entities/auth/api'
import { Button, Input } from '@nextui-org/react'
import { useFormik } from 'formik'
import * as Yup from 'yup'

type SignUpForm = {
	username: string
	password: string
	repeatPassword: string
}

const validationSchema = Yup.object().shape({
	username: Yup.string().required('Логин обязателен'),
	password: Yup.string().required('Пароль обязателен'),
	repeatPassword: Yup.string().oneOf([Yup.ref('password')], 'Пароли должны совпадать'),
})

export const SignUpForm = () => {
	const formik = useFormik<SignUpForm>({
		initialValues: {
			username: '',
			password: '',
			repeatPassword: '',
		},
		initialErrors: {
			username: '',
			password: '',
			repeatPassword: '',
		},
		validateOnChange: true,
        validateOnBlur: true,
        validateOnMount: false,
		validationSchema, 
		onSubmit: (values) => {
			alert(JSON.stringify(values, null, 2))

			// const data = {
			// 	username: values.username,
			// 	password: values.password
			// }

			// signUp(data)
		},
	})

	return (
		<>
			<form onSubmit={formik.handleSubmit} className='flex flex-col gap-4'>
				<Input
					type='text'
					name='username'
					placeholder='Логин'
					isInvalid={!!formik.errors.username}
					value={formik.values.username}
					errorMessage={formik.errors.username}
					onChange={formik.handleChange}
				/>
				<Input
					type='password'
					name='password'
					placeholder='Пароль'
					isInvalid={!!formik.errors.password}
					value={formik.values.password}
					errorMessage={formik.errors.password}
					onChange={formik.handleChange}
				/>
				<Input
					type='password'
					name='repeatPassword'
					placeholder='Повторите пароль'
					isInvalid={!!formik.errors.repeatPassword}
					value={formik.values.repeatPassword}
					errorMessage={formik.errors.repeatPassword}
					onChange={formik.handleChange}
				/>
				<Button type='submit' color='primary'>
					Зарегистрироваться
				</Button> 
			</form>
		</>
	)
}
