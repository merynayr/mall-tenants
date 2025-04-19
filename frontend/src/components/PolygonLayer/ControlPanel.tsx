import styles from './ControlPanel.module.css';

interface ControlPanelProps {
  isDrawing: boolean;
  onStart: () => void;
  onFinish: () => void;
  onReset: () => void;
  floor: number;
  onFloorChange: (val: number) => void;
}

export const ControlPanel: React.FC<ControlPanelProps> = ({
	isDrawing,
	onStart,
	onFinish,
	onReset,
	floor,
	onFloorChange
}) => (
	<div className={styles.panel}>
		<div className={styles.topRow}>
			<button onClick={onStart} disabled={isDrawing}>Начать рисовать</button>
			<button onClick={onFinish} disabled={!isDrawing}>Завершить фигуру</button>
			<button onClick={onReset}>Сбросить</button>
			<div className={styles.floorInput}>
				<label>Этаж:</label>
				<input
				 type="number"
				 value={floor}
				 onChange={(e) => onFloorChange(Number(e.target.value))} 
				/>
			</div>
		</div>

	</div>
);
