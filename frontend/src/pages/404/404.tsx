import styles from './404.module.css';

const Page404: React.FC = () => {
	return (
		<div className={styles.container}>
			<p className={styles.code}>404</p>
			<h1 className={styles.heading}>Страница не найдена</h1>
			<p className={styles.text}>
				Скорее всего, ты тут из-за очепятки в адресе страницы.
			</p>
			<p className={styles.text}>
				Попробуй <a href="/" className={styles.link}>вернуться на главную</a> или свяжись с администрацией сайта.
			</p>
			<img
				src="/public/404.png"
				srcSet="/images/static/404@2x.jpg 2x"
				alt="404"
				className={styles.image}
			/>
		</div>
	);
};

export default Page404;
