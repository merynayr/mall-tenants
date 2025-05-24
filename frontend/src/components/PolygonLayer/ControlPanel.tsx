import cn from 'classnames';
import styles from './ControlPanel.module.css';
import { useHasRole } from '@/hooks/Role';

interface ControlPanelProps {
  isDrawing: boolean;
  onStart: () => void;
  onFinish: () => void;
  onReset: () => void;
  floor: number;
  onFloorChange: (val: number) => void;
	DeletePolygon: (code: number | null) => void;
	selectedPolygonIndex: number | null;
	onAddFloorPlan: () => void;
}

export const ControlPanel: React.FC<ControlPanelProps> = ({
	isDrawing,
	onStart,
	onFinish,
	onReset,
	floor,
	onFloorChange,
	DeletePolygon,
	selectedPolygonIndex,
	onAddFloorPlan
}) => (
	<div className={styles.wrapperRow}>
		<div className={styles.panel}>
			<div className={styles.topRow}>
				<div className={styles.floorInput}>
					<label>Этаж:</label>
					<input
						className={styles.input}
						type="number"
						value={floor}
						onChange={(e) => onFloorChange(Number(e.target.value))} />
				</div>

				{useHasRole('moderator', 'director') && (
					<>
						<div className={styles.verticalDivider} />
						<div className={styles.buttonGroup}>
							<button onClick={onStart} disabled={isDrawing}>Начать рисовать</button>
							<button onClick={onFinish} disabled={!isDrawing}>Завершить фигуру</button>
							<button onClick={onReset} disabled={!isDrawing}>Сбросить</button>
							<button
								onClick={() => DeletePolygon(selectedPolygonIndex)}
								disabled={selectedPolygonIndex === null}
							>
							Удалить выделенное
							</button>
						</div>
					</>
				)}
			</div>
		</div>

		{useHasRole('moderator', 'director') && (
			<div className={cn(styles.rightAligned, styles.buttonGroup)}>
				<button onClick={onAddFloorPlan}>Добавить новый план этажа</button>
			</div>
		)}
	</div>
);
