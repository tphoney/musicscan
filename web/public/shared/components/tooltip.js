import styles from "./tooltip.module.css";
import classnames from "classnames";

import * as Tooltip from "@radix-ui/react-tooltip";

export default (props) => (
	<Tooltip.Root>
		<Tooltip.Trigger className={styles.trigger}>
			{props.children}
		</Tooltip.Trigger>
		<Tooltip.Content className={classnames(styles.root, props.className)}>
			{props.content}
		</Tooltip.Content>
	</Tooltip.Root>
);
