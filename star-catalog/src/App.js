import React, { useState } from 'react'
import './App.css'
import SearchForm from './components/SearchForm'
import AddStarForm from './components/AddStarForm'

export default function App() {
	const [results, setResults] = useState([])
	const [message, setMessage] = useState('')

	const handleSearch = async params => {
		setMessage('')
		setResults([])

		try {
			const url = new URL('http://localhost:3000/api/stars')
			Object.entries(params).forEach(([key, value]) => {
				if (value !== undefined && value !== null && value !== '') {
					url.searchParams.append(key, value)
				}
			})

			const response = await fetch(url)
			if (!response.ok) {
				const errorText = await response.text()
				throw new Error(errorText || `HTTP ${response.status}`)
			}

			const data = await response.json()
			const starsArray = Array.isArray(data) ? data : data ? [data] : []
			setResults(starsArray)
		} catch (err) {
			setMessage('❌ Ошибка поиска: ' + (err.message || 'неизвестная ошибка'))
		}
	}

	const handleAddStar = async star => {
		setMessage('')
		try {
			const response = await fetch('http://localhost:3000/api/stars', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(star),
			})

			if (!response.ok) {
				const errorText = await response.text()
				throw new Error(errorText || `HTTP ${response.status}`)
			}

			const savedStar = await response.json()
			setMessage(`✅ Звезда "${savedStar.name}" успешно добавлена!`)
			setResults([]) // очищаем результаты поиска
		} catch (err) {
			setMessage(
				'❌ Ошибка добавления: ' + (err.message || 'неизвестная ошибка')
			)
		}
	}

	return (
		<div>
			<header>
				<h1>🔭 Каталог звёзд нашей галактики</h1>
				<p>Исследуйте звёзды по названию, созвездию, расстоянию и массе</p>
			</header>

			<main>
				<SearchForm onSearch={handleSearch} />
				<AddStarForm onAdd={handleAddStar} />

				<div className='results'>
					{message && <div className='success-message'>{message}</div>}
					{results.length > 0 ? (
						<>
							<h2>Найдено звёзд: {results.length}</h2>
							<table>
								<thead>
									<tr>
										<th>Название</th>
										<th>Созвездие</th>
										<th>Расстояние (св. лет)</th>
										<th>
											Масса (M<sub>☉</sub>)
										</th>
									</tr>
								</thead>
								<tbody>
									{results.map((star, i) => (
										<tr key={i}>
											<td>{star.name || ''}</td>
											<td>{star.constellation || ''}</td>
											<td>
												{star.distance !== undefined ? star.distance : ''}
											</td>
											<td>{star.mass !== undefined ? star.mass : ''}</td>
										</tr>
									))}
								</tbody>
							</table>
						</>
					) : (
						!message && (
							<p className='hint'>
								Введите параметры поиска и нажмите «Найти звёзды»
							</p>
						)
					)}
				</div>
			</main>

			<footer>
				<p>
					Проект по дисциплине «Разработка приложений с базами данных» | Вариант
					5
				</p>
			</footer>
		</div>
	)
}
