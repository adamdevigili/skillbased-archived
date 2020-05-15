import SportsDropdown from '../components/Dropdown';
import GoButton from '../components/GoButton';
import Title from '../components/Title';

export default () => (
	<div className="centered">
		<div className="title">
			<h1>SKILLBASED.IO</h1>
		</div>
		<div className="sport_dropdown">
			<h3>Select a sport to get started</h3>
			<SportsDropdown />
			<div className="go_button">
				<GoButton />
			</div>
		</div>
	</div>
);
