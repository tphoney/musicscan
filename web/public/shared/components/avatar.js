import styles from "./avatar.module.css";
import classnames from "classnames";

import * as Avatar from "@radix-ui/react-avatar";

export default (props) => (
	<Avatar.Root className={classnames(styles.root, props.className)}>
		<Avatar.Image className={styles.image} src={props.src} />
		<Avatar.Fallback className={styles.fallback}>
			{props.text && props.text.charAt(0)}
		</Avatar.Fallback>
	</Avatar.Root>
);
