const URL = import.meta.env.VITE_URL;

export const postBugData = async (bugText: string) => {
	let url = `${URL}/submitbug`;

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			bugText: bugText,
		}),
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.ok;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
