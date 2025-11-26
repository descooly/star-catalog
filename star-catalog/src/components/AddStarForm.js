import React, { useState } from 'react'

export default function AddStarForm({ onAdd }) {
	const [formData, setFormData] = useState({
		name: '',
		constellation: '',
		distance: '',
		mass: '',
	})

	const handleChange = e => {
		setFormData({ ...formData, [e.target.name]: e.target.value })
	}

	const handleSubmit = e => {
		e.preventDefault()
		const star = {
			name: formData.name.trim(),
			constellation: formData.constellation.trim(),
			distance: formData.distance ? parseFloat(formData.distance) : undefined,
			mass: formData.mass ? parseFloat(formData.mass) : undefined,
		}

		// Проверка обязательных полей
		if (
			!star.name ||
			!star.constellation ||
			star.distance === undefined ||
			star.mass === undefined
		) {
			alert('Все поля обязательны для заполнения')
			return
		}

		onAdd(star)
	}

	return (
		<div className='add-star-box'>
			<h2>✨ Добавить новую звезду</h2>
			<form onSubmit={handleSubmit}>
				<div className='form-group'>
					<label htmlFor='newStarName'>Название звезды:</label>
					<input
						type='text'
						id='newStarName'
						name='name'
						value={formData.name}
						onChange={handleChange}
						required
						placeholder='Например: Вега'
					/>
				</div>

				<div className='form-group'>
					<label htmlFor='newConstellation'>Созвездие:</label>
					<input
						type='text'
						id='newConstellation'
						name='constellation'
						value={formData.constellation}
						onChange={handleChange}
						required
						placeholder='Например: Лира'
					/>
				</div>

				<div className='form-row'>
					<div className='form-group'>
						<label htmlFor='newDistance'>Расстояние от Солнца (св. лет):</label>
						<input
							type='number'
							id='newDistance'
							name='distance'
							value={formData.distance}
							onChange={handleChange}
							required
							step='0.1'
						/>
					</div>
					<div className='form-group'>
						<label htmlFor='newMass'>Масса (в массах Солнца):</label>
						<input
							type='number'
							id='newMass'
							name='mass'
							value={formData.mass}
							onChange={handleChange}
							required
							step='0.01'
						/>
					</div>
				</div>

				<button type='submit'>➕ Добавить звезду</button>
			</form>
		</div>
	)
}
