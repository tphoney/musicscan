import styles from "./demo.module.css";

import Avatar from "../avatar";
import Button from "../button";
import Breadcrumb from "../breadcrumb";
import Checkbox from "../checkbox";
import Input from "../input";
import Menu from "../menu";
import Search from "../search";
import Select from "../select";
import Switch from "../switch";
import Tooltip from "../tooltip";
import Textarea from "../textarea";

export default (props) => {
	return (
		<div className={styles.root}>
			<section>
				<h2>Button</h2>
				<div>
					<Button
						onClick={() => console.log("clicked")}
						onMouseEnter={() => console.log("enter")}
						onMouseLeave={() => console.log("leave")}
					>
						Button
					</Button>
				</div>
				<div>
					<Button disabled onClick={() => console.log("clicked")}>
						Disabled
					</Button>
				</div>
			</section>
			<section>
				<h2>Text Input</h2>
				<div>
					<Input />
				</div>
				<div>
					<Input disabled />
				</div>
				<div>
					<Input type="password" placeholder="Password" />
				</div>
			</section>
			<section>
				<h2>Search Input</h2>
				<div>
					<Search />
				</div>
				<div>
					<Search placeholder="Search Projects …" />
				</div>
				<div>
					<Search disabled />
				</div>
			</section>
			<section>
				<h2>Search Input</h2>
				<div>
					<Select onChange={() => console.log("changed")}>
						<option>Apple</option>
						<option>Orange</option>
					</Select>
				</div>
				<div>
					<Select disabled>
						<option>Apple</option>
						<option>Orange</option>
					</Select>
				</div>
			</section>
			<section>
				<h2>Textarea</h2>
				<div>
					<Textarea></Textarea>
				</div>
				<div>
					<Textarea disabled />
				</div>
			</section>
			<section>
				<h2>Menu</h2>
				<div>
					<Menu>
						<Menu.Trigger>
							<Button>Click</Button>
						</Menu.Trigger>
						<Menu.Content>
							<Menu.Item>Cut</Menu.Item>
							<Menu.Item>Copy</Menu.Item>
							<Menu.Item>Paste</Menu.Item>
						</Menu.Content>
					</Menu>
				</div>
			</section>
			<section>
				<h2>Tooltip</h2>
				<div>
					<Tooltip content="Hello World">
						<Button>Hover</Button>
					</Tooltip>
				</div>
			</section>
			<section>
				<h2>Avatar</h2>
				<div>
					<Avatar src="https://avatars.githubusercontent.com/u/817538" />
				</div>
				<div>
					<Avatar text="Brad" />
				</div>
			</section>
			<section>
				<h2>Breadcrumb</h2>
				<div>
					<Breadcrumb>
						<a href="#">Components</a>
						<a href="#">Core</a>
						<a href="#">Breadcrumb</a>
					</Breadcrumb>
				</div>
			</section>
			<section>
				<h2>Checkbox</h2>
				<div>
					<Checkbox onCheckedChange={() => console.log("changed")}></Checkbox>
				</div>
				<div>
					<Checkbox disabled={true}></Checkbox> Disabled
				</div>
			</section>
			<section>
				<h2>Dialog</h2>
			</section>
			<section>
				<h2>Switch</h2>
				<div>
					<Switch onCheckedChange={() => console.log("changed")}></Switch>
				</div>
				<div>
					<Switch checked={true}></Switch> Checked
				</div>
				<div>
					<Switch disabled={true}></Switch> Disabled
				</div>
			</section>
		</div>
	);
};
