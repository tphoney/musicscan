import styles from "./menu.module.css";
import classnames from "classnames";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";

const Menu = (props) => {
	return <DropdownMenu.Root>{props.children}</DropdownMenu.Root>;
};

Menu.Item = (props) => (
	<DropdownMenu.Item className={classnames(styles.item, props.className)}>
		{props.children}
	</DropdownMenu.Item>
);

Menu.Content = (props) => (
	<DropdownMenu.Content className={classnames(styles.content, props.className)}>
		{props.children}
	</DropdownMenu.Content>
);

Menu.Trigger = (props) => (
	<DropdownMenu.Trigger className={classnames(styles.trigger, props.className)}>
		{props.children}
	</DropdownMenu.Trigger>
);

export default Menu;
