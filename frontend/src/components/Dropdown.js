import React from 'react';
import { Dropdown } from 'semantic-ui-react';

const sportOptions = [
	{text: 'Ultimate Frisbee' },
	{text: 'Basketball' },
	{text: 'Football' }
];

const SportsDropdown = () => (
	<Dropdown placeholder="Select Sport" fluid search selection options={sportOptions} />
);

export default SportsDropdown;
