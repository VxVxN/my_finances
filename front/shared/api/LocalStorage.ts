const getTokensLocalStorage = () => {
	const accessToken = window.localStorage.getItem('accessToken')
	const refreshToken = window.localStorage.getItem('refreshToken')

	return { accessToken, refreshToken }
}
