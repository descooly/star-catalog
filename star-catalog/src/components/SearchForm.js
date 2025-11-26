import React, { useState } from 'react'

export default function SearchForm({ onSearch }) {
	const [formData, setFormData] = useState({
		starName: '',
		constellation: '',
		minDistance: '',
		maxDistance: '',
		minMass: '',
		maxMass: '',
	})

	const handleChange = e => {
		setFormData({ ...formData, [e.target.name]: e.target.value })
	}

	const handleSubmit = e => {
		e.preventDefault()
		const params = {
			starName: formData.starName.trim() || undefined,
			constellation: formData.constellation.trim() || undefined,
			minDistance: formData.minDistance || undefined,
			maxDistance: formData.maxDistance || undefined,
			minMass: formData.minMass || undefined,
			maxMass: formData.maxMass || undefined,
		}
		onSearch(params)
	}

	return (
		<div className='search-box'>
			<h2>🔍 Найти звезду</h2>
			<form onSubmit={handleSubmit}>
				<div className='form-group'>
					<label htmlFor='starName'>Название звезды:</label>
					<input
						type='text'
						id='starName'
						name='starName'
						value={formData.starName}
						onChange={handleChange}
						placeholder='Например: Сириус'
					/>
				</div>

				<div className='form-group'>
					<label htmlFor='constellation'>Созвездие:</label>
					<input
						type='text'
						id='constellation'
						name='constellation'
						value={formData.constellation}
						onChange={handleChange}
						placeholder='Например: Большой Пёс'
					/>
				</div>

				<div className='form-row'>
					<div className='form-group'>
						<label htmlFor='minDistance'>
							Расстояние от Солнца, св. лет (от):
						</label>
						<input
							type='number'
							id='minDistance'
							name='minDistance'
							value={formData.minDistance}
							onChange={handleChange}
							step='0.1'
						/>
					</div>
					<div className='form-group'>
						<label htmlFor='maxDistance'>до:</label>
						<input
							type='number'
							id='maxDistance'
							name='maxDistance'
							value={formData.maxDistance}
							onChange={handleChange}
							step='0.1'
						/>
					</div>
				</div>

				<div className='form-row'>
					<div className='form-group'>
						<label htmlFor='minMass'>Масса (в массах Солнца, от):</label>
						<input
							type='number'
							id='minMass'
							name='minMass'
							value={formData.minMass}
							onChange={handleChange}
							step='0.01'
						/>
					</div>
					<div className='form-group'>
						<label htmlFor='maxMass'>до:</label>
						<input
							type='number'
							id='maxMass'
							name='maxMass'
							value={formData.maxMass}
							onChange={handleChange}
							step='0.01'
						/>
					</div>
				</div>

				<button type='submit'>🔍 Найти звёзды</button>
			</form>
		</div>
	)
}
