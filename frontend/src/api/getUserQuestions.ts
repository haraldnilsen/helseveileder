interface Question {
	QuestionID: number;
	QuestionText: string;
}

interface Answer {
	AnswerID: number;
	QuestionID: number;
	IsChatGPT: boolean;
	AnswerText: string;
}

interface QuestionAnswerPair {
	Question: Question;
	Answers: Answer[];
}

interface QAData {
	questions: QuestionAnswerPair[];
}

export const getUserQuestions = (respondentID: number): Promise<QAData> => {
	let url = `http://localhost:8080/userquestions?respondentID=${respondentID}`;

	const response = fetch(url, {
		method: "GET",
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.json();
		})
		.then((data) => {
			console.log(data);
			localStorage.setItem("userQuestions", JSON.stringify(data));
			return data;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
