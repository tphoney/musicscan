import styles from "./switch.module.css";
import classnames from "classnames";

import * as Switch from "@radix-ui/react-switch";

export default (props) => (
	<Switch.Root
		className={classnames(styles.root, props.className)}
		onCheckedChange={props.onCheckedChange}
		checked={props.checked}
		disabled={props.disabled}
	>
		<Switch.Thumb className={styles.thumb} />
	</Switch.Root>
);
