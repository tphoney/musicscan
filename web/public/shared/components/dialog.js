import styles from "./dialog.module.css";
import classnames from "classnames";

import * as Dialog from "@radix-ui/react-dialog";

export default () => (
	<Dialog.Root>
		<Dialog.Trigger />
		<Dialog.Overlay />
		<Dialog.Content>
			<Dialog.Close />
		</Dialog.Content>
	</Dialog.Root>
);
