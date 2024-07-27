'use client'

import { Button, Input } from '@nextui-org/react'
import { useFormik } from 'formik'
import * as Yup from 'yup'

type SignInForm = {
	username: string
	password: string
}

const validationSchema = Yup.object().shape({
	username: Yup.string().required('Логин обязателен'),
	password: Yup.string().required('Пароль обязателен'),
})

const SignInForm = () => {
	const formik = useFormik<SignInForm>({
		initialValues: {
			username: 'vxvxn',
			password: '123',
		},
		validateOnChange: true,
		validateOnBlur: true,
		validateOnMount: false,
		validationSchema,
		onSubmit: values => {
			const data = {
				username: values.username,
				password: values.password,
			}

			// signIn(data).then((response) => {
			//   console.log(response);
			// });

			fetch('http://localhost:8080/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(data),
			}).then(response => {
				console.log(response)
			})
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
				<Button type='submit' color='primary'>
					Войти
				</Button>
			</form>
		</>
	)
}

export default SignInForm
