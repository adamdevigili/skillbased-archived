import React from 'react';
import { Dropdown } from 'semantic-ui-react';

const sportOptions = [
	{ key: 'uf', value: 'uf', text: 'Ultimate Frisbee' },
	{ key: 'bb', value: 'bb', text: 'Basketball' },
	{ key: 'ff', value: 'ff', text: 'Football' }
];

const SportsDropdown = () => (
	<Dropdown placeholder="Select Sport" fluid search selection options={sportOptions} />
);

export default SportsDropdown;
