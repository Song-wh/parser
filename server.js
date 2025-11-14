const http = require('http');
const fs = require('fs');
const path = require('path');

const PORT = 9990;

// MIME 타입 매핑
const mimeTypes = {
  '.html': 'text/html',
  '.js': 'text/javascript',
  '.css': 'text/css',
  '.json': 'application/json',
  '.png': 'image/png',
  '.jpg': 'image/jpg',
  '.gif': 'image/gif',
  '.svg': 'image/svg+xml',
  '.ico': 'image/x-icon'
};

const server = http.createServer((req, res) => {
  console.log(`${req.method} ${req.url}`);

  // 루트 경로는 index.html로 처리
  let filePath = req.url === '/' ? '/index.html' : req.url;
  filePath = path.join(__dirname, filePath);

  const extname = String(path.extname(filePath)).toLowerCase();
  const contentType = mimeTypes[extname] || 'application/octet-stream';

  fs.readFile(filePath, (error, content) => {
    if (error) {
      if (error.code === 'ENOENT') {
        res.writeHead(404, { 'Content-Type': 'text/html; charset=utf-8' });
        res.end('<h1>404 - 파일을 찾을 수 없습니다</h1>', 'utf-8');
      } else {
        res.writeHead(500);
        res.end(`서버 오류: ${error.code}`, 'utf-8');
      }
    } else {
      res.writeHead(200, { 
        'Content-Type': contentType,
        'Access-Control-Allow-Origin': '*',
        'Cache-Control': 'no-store, no-cache, must-revalidate'
      });
      res.end(content, 'utf-8');
    }
  });
});

server.listen(PORT, () => {
  console.log('\n========================================');
  console.log('🚀 서버가 시작되었습니다!');
  console.log(`📡 http://localhost:${PORT}`);
  console.log('========================================');
  console.log('종료하려면 Ctrl+C를 누르세요.\n');
});

// 에러 핸들링
server.on('error', (err) => {
  if (err.code === 'EADDRINUSE') {
    console.error(`❌ 포트 ${PORT}가 이미 사용 중입니다.`);
    process.exit(1);
  } else {
    console.error('서버 에러:', err);
  }
});

