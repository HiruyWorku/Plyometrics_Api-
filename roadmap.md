# 🏋️ Jump Analyzer Roadmap

## Phase 1 – Foundations (Week 1–2)
- Research existing jump analysis tools (e.g., MyJump2).
- Choose stack: **Go (Gin API)** + **Python (MCP + pose detection)**.
- Set up repo: `backend/` (Go), `analysis/` (Python MCP), `client/` (optional web UI).
- Build Go API `/upload-video` → stores video + calls MCP.
- Create MCP server in Python → dummy tool returns placeholder results.

---

## Phase 2 – Core Jump Height Estimator (Week 3–4)
- Implement pose detection with **MediaPipe** or **MoveNet**.
- Track hip/knee/ankle positions across frames.
- Compute vertical displacement → estimate jump height.
- Scale pixel measurements → real-world estimate (based on user height input).
- Expose MCP tool:
  ```python
  @mcp.tool
  def estimate_jump_height(video: bytes, user_height: float) -> float:
      ...
  ```
- Connect Go API to call MCP tool → return height JSON.

---

## Phase 3 – Form Feedback System (Week 5–6)
- Implement rule-based heuristics:
  - Depth of countermovement.
  - Knee alignment (valgus detection).
  - Arm swing contribution.
- Return structured feedback:
  ```json
  {
    "height": 24.5,
    "feedback": "Good depth, but knees collapsing inward."
  }
  ```
- Add second MCP tool:
  ```python
  @mcp.tool
  def analyze_jump_form(video: bytes) -> dict:
      ...
  ```

---

## Phase 4 – App Layer (Week 7–8)
- Build simple web client (React or HTML).
- Upload video → results page.
- Add user accounts + history (Postgres/SQLite).
- Graph jump height progression.

---

## Phase 5 – Monetization + Polish (Week 9–10)
- Add **Freemium Model**:
  - Free: Jump height only.
  - Paid: Detailed form analysis + history tracking.
- Deploy Go backend (Railway/Render).
- Package MCP analysis in Docker.
- Optimize model performance (downsample video, GPU acceleration).

---

## Stretch Goals
- Real-time form feedback (live camera stream).
- Team dashboard for coaches.
- Personalized training recommendations (AI-based).
