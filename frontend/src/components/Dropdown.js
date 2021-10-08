import React, {Fragment, useState, useEffect} from 'react';
// import { Dropdown } from 'semantic-ui-react';
import { Select } from 'antd';
import http from 'http';
import axios from 'axios';

let sportOptions = [];

const { Option } = Select;

function handleChange(value) {
	console.log(`Selected: ${value}`);
}

function SportsDropdown() {
	const [isLoading, setIsLoading] = useState(false);
	const apiURL = process.env.REACT_APP_SKILLBASED_API_URL + process.env.REACT_APP_SKILLBASED_API_VERSION

	useEffect(() => {
		const fetchData = async () => {
			setIsLoading(true);


			const result = await axios(
				apiURL + '/sports',
			);

			result.data.items.forEach((sport) => {
				sportOptions.push(<Option key={sport.name}>{sport.name}</Option>)
			})
			setIsLoading(false);
		};

		fetchData()
	}, []);

	return (
		<div>
			{isLoading ? (
				<Select
					placeholder="Loading"
					size="large"
					style={{
						width: '100%',
						margin: 'auto'
					}}
					showSearch
					fluid
					search
					selection
					loading
					disabled
				/>
			) : (
				<Select
					placeholder="Select Sport"
					size="large"
					style={{
						width: '100%',
						margin: 'auto',
						cursor: 'pointer'
					}}
					onChange={handleChange}
					showSearch
					fluid
					search
					selection
				>
					{sportOptions}
				</Select>
			)}
		</div>
	)
}

export default SportsDropdown;
