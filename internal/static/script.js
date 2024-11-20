document.addEventListener("DOMContentLoaded", () => {
    console.log("DOM fully loaded and parsed.");
    // 현재 시간 폼 처리
    const currentTimeForm = document.getElementById("current-time-form");
    const currentTimeResult = document.getElementById("current-time-result");

    let currentInterval; // 현재 시간 타이머를 저장하기 위한 변수

    currentTimeForm.addEventListener("submit", async (event) => {
        event.preventDefault(); // 폼 제출 시 페이지 새로고침 방지
        const timezone = document.getElementById("timezone").value.trim(); // 입력된 타임존 가져오기
        if (!timezone) {
            currentTimeResult.textContent = "유효한 타임존을 입력해주세요."; // 타임존이 비어있으면 에러 메시지 출력
            return;
        }

        try {
            // 서버에 요청하여 현재 시간을 가져옴
            const response = await fetch(`/api?timezone=${encodeURIComponent(timezone)}`);
            const data = await response.json();
            if (response.ok) {
                if (currentInterval) clearInterval(currentInterval); // 이전 타이머가 있으면 중지
                let currentTime = new Date(data.time); // 서버에서 받은 시간을 Date 객체로 변환

                // 화면에 시간 출력
                currentTimeResult.textContent = `현재 ${data.timezone} 시간: ${currentTime.toLocaleTimeString()}`;

                // 매초 시간 업데이트
                currentInterval = setInterval(() => {
                    currentTime.setSeconds(currentTime.getSeconds() + 1); // 1초 증가
                    currentTimeResult.textContent = `현재 ${data.timezone} 시간: ${currentTime.toLocaleTimeString()}`;
                }, 1000); // 1초마다 실행
            } else {
                currentTimeResult.textContent = data.error; // 서버에서 에러 메시지가 있으면 출력
            }
        } catch (error) {
            currentTimeResult.textContent = "오류가 발생했습니다. 다시 시도해주세요."; // 요청 실패 시 에러 메시지 출력
        }
    });

    // 세계 시간 정보를 불러오는 버튼 처리
    const loadWorldTimesButton = document.getElementById("load-world-times");
    const worldTimesResult = document.getElementById("world-times-result");

    loadWorldTimesButton.addEventListener("click", async () => {
        try {
            // 서버에서 세계 시간 데이터를 요청
            const response = await fetch("/api/world");
            const data = await response.json();
            if (response.ok) {
                worldTimesResult.innerHTML = ""; // 기존 내용을 초기화
                for (const [city, time] of Object.entries(data.city_times)) {
                    const div = document.createElement("div"); // 도시별 시간을 표시할 div 생성
                    div.className = "city-time";
                    div.innerHTML = `<h3>${city}</h3><p>${time}</p>`;
                    worldTimesResult.appendChild(div); // 결과 영역에 추가

                    // 각 도시의 시간을 업데이트하는 타이머 설정
                    let cityTime = new Date(time);
                    setInterval(() => {
                        cityTime.setSeconds(cityTime.getSeconds() + 1); // 1초 증가
                        div.querySelector("p").textContent = cityTime.toLocaleTimeString(); // 화면 업데이트
                    }, 1000); // 1초마다 실행
                }
            } else {
                worldTimesResult.textContent = data.error; // 서버 에러 메시지 출력
            }
        } catch (error) {
            worldTimesResult.textContent = "오류가 발생했습니다. 다시 시도해주세요."; // 요청 실패 시 에러 메시지 출력
        }
    });
});
