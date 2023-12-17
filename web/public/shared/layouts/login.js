import styles from "./login.module.css";
import classnames from "classnames";

export default (props) => (
	<div className={classnames(styles.root, props.className)}>
		{props.children}
	</div>
);
