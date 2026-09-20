# Port Detective — Task Tracker

> Cập nhật file này mỗi khi hoàn thành hoặc bắt đầu một task.
> Quy ước: `[ ]` chưa làm · `[/]` đang làm · `[x]` đã xong

---

## Phase 1: Foundation

- [x] **Task 1:** Khởi tạo Go module + cấu trúc thư mục
  - `go mod init github.com/Khoa180806/Port_Detective`
  - Tạo thư mục `cmd/`, `internal/lookup/`, `internal/process/`, `internal/output/`
  - `.gitignore` cho Go project
  - `main.go` placeholder — `go build ./...` thành công
  - Kết nối remote GitHub

- [ ] **Task 2:** Data model — `ProcessInfo` struct
  - Struct `ProcessInfo{PID, Name, Command, Port, Protocol}` với JSON tags
  - Unit test kiểm tra JSON marshal/unmarshal

- [ ] **Task 3:** Output formatter — Text + JSON
  - `FormatText(*ProcessInfo)` → output dạng bảng đẹp
  - `FormatJSON(*ProcessInfo)` → JSON valid
  - `FormatTextMultiple` / `FormatJSONMultiple` cho danh sách
  - Unit test cho cả 2 formatter

### ✅ Checkpoint: Foundation
- [ ] `go build ./...` thành công
- [ ] `go test ./...` tất cả pass

---

## Phase 2: Core Lookup

- [ ] **Task 4:** Interface `PortLookupStrategy` + dispatch
  - `FindProcessByPort(port int) ([]ProcessInfo, error)`
  - `KillProcess(pid int) error`

- [ ] **Task 5:** Windows implementation — `netstat` + `tasklist`
  - Parse `netstat -ano` → PID theo port
  - Parse `tasklist` / `wmic` → tên process + command line
  - Unit test với sample output
  - Test thủ công trên port thật

- [ ] **Task 6:** Linux implementation — `lsof` / `/proc/net/tcp`
  - Parse `lsof -i :<port>` output
  - Fallback parse `/proc/net/tcp`
  - Unit test với sample output

- [ ] **Task 7:** macOS (Darwin) implementation — `lsof`
  - Parse `lsof -i :<port>` output cho macOS
  - Unit test với sample output

### ✅ Checkpoint: Core Lookup
- [ ] `go build ./...` thành công
- [ ] `go test ./...` tất cả pass
- [ ] Chạy thử `FindProcessByPort` trên Windows với port thật

---

## Phase 3: CLI Commands (Cobra)

- [ ] **Task 8:** Cài Cobra + Root command
  - `go get github.com/spf13/cobra`
  - Root command với description, version, help tự động
  - `port-detective --help` hoạt động

- [ ] **Task 9:** `check` command
  - `port-detective check <port>` → hiển thị process info
  - Flag `--json` cho JSON output
  - Exit code: 0=tìm thấy, 1=port free, 2=lỗi
  - Validate port input (1–65535)

- [ ] **Task 10:** `kill` command
  - `port-detective kill <port>` → kill process
  - Prompt xác nhận y/n (default=no)
  - Flag `--force` bỏ qua prompt
  - Flag `--json` cho JSON output
  - Xử lý edge case: nhiều process cùng port

### ✅ Checkpoint: CLI Commands
- [ ] `port-detective check <port>` hoạt động end-to-end
- [ ] `port-detective kill <port>` hoạt động end-to-end
- [ ] `--json` flag hoạt động cho cả 2 command

---

## Phase 4: Polish & Cross-platform

- [ ] **Task 11:** `fatih/color` + beautify text output
  - Output text có màu (PID xanh, process vàng, tiêu đề đỏ)
  - Tắt màu tự động khi pipe/redirect

- [ ] **Task 12:** Cross-compile check
  - `GOOS=linux go build` thành công
  - `GOOS=darwin go build` thành công
  - `GOOS=windows go build` thành công

- [ ] **Task 13:** Error handling toàn diện
  - Custom error types
  - Thông báo thân thiện cho từng loại lỗi
  - Exit code nhất quán

### ✅ Checkpoint: Polish
- [ ] Cross-compile 3 OS thành công
- [ ] Error handling thân thiện
- [ ] Output đẹp có màu

---

## Phase 5: Documentation & Release

- [ ] **Task 14:** README.md hoàn chỉnh
  - Mô tả, installation, usage examples, demo output

- [ ] **Task 15:** `scan` command (Optional — M5)
  - `port-detective scan <start>-<end>` quét dải port
  - `--json` output JSON array
  - Hiển thị progress

- [ ] **Task 16:** GoReleaser config + push final
  - `.goreleaser.yml` cho 3 OS × 2 arch
  - Push tất cả lên GitHub

### ✅ Checkpoint: Final
- [ ] `go build ./...` + `go test ./...` pass
- [ ] Cross-compile 3 OS thành công
- [ ] README đầy đủ
- [ ] Code đã push GitHub
