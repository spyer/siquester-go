// SIQuester Web App
const API_BASE = '/api';

// Translations
const i18n = {
    en: {
        // Header
        new: 'New',
        open: 'Open',
        save: 'Save',
        export: 'Export',
        exportSiq: 'Export as .siq',
        exportXml: 'Export as XML',
        exportYaml: 'Export as YAML',
        toggleTheme: 'Toggle theme',
        language: 'Language',

        // Sidebar
        packages: 'Packages',
        refresh: 'Refresh',
        noPackagesOpen: 'No packages open',
        createNew: 'Create New',
        close: 'Close',

        // Welcome
        welcomeTitle: 'Welcome to SIQuester',
        welcomeSubtitle: 'A cross-platform editor for SIGame question packages',
        createNewPackage: 'Create New Package',
        openExistingPackage: 'Open Existing Package',

        // Tree
        packageStructure: 'Package Structure',
        addRound: 'Add Round',
        addTheme: 'Add Theme',
        addQuestion: 'Add Question',
        delete: 'Delete',
        media: 'Media',
        files: 'files',
        selectItemToEdit: 'Select an item from the tree to edit',
        noPackageLoaded: 'No package loaded',
        empty: '(empty)',

        // Package Editor
        packageSettings: 'Package Settings',
        general: 'General',
        name: 'Name',
        date: 'Date',
        packageLanguage: 'Language',
        publisher: 'Publisher',
        difficulty: 'Difficulty',
        restriction: 'Restriction',
        restrictionPlaceholder: 'e.g., 18+',
        authors: 'Authors',
        addAuthor: '+ Add Author',
        comments: 'Comments',

        // Media Manager
        mediaManager: 'Media Manager',
        images: 'Images',
        audio: 'Audio',
        video: 'Video',
        html: 'HTML',
        uploadNewMedia: 'Upload New Media',
        noMediaFiles: 'No {type} files',

        // Round Editor
        round: 'Round',
        type: 'Type',
        roundTypeStandard: 'Standard',
        roundTypeFinal: 'Final',

        // Theme Editor
        theme: 'Theme',

        // Question Editor
        questionEditor: 'Question Editor',
        price: 'Price',
        questionTypeDefault: 'Default',
        questionTypeStake: 'Stake (Auction)',
        questionTypeSecret: 'Secret (Cat in a Bag)',
        questionTypeNoRisk: 'No Risk',
        questionTypeForAll: 'For All (Simultaneous)',
        questionContent: 'Question Content',
        text: 'Text',
        questionText: 'Question Text',
        rightAnswers: 'Right Answers',
        addAnswer: '+ Add Answer',
        wrongAnswers: 'Wrong Answers (Optional)',
        addWrongAnswer: '+ Add Wrong Answer',
        addText: '+ Text',
        addMedia: '+ Media',
        mediaAdded: 'Media added to question',
        noMediaUploaded: 'No media uploaded yet. Upload media first.',
        answerContent: 'Answer Content (displayed when revealing answer)',
        answerText: 'Answer Text',
        mediaAddedToAnswer: 'Media added to answer',

        // Toasts
        packageCreated: 'Package created',
        packageOpened: 'Package opened',
        packageClosed: 'Package closed',
        roundAdded: 'Round added',
        roundDeleted: 'Round deleted',
        roundMoved: 'Round moved',
        themeAdded: 'Theme added',
        themeDeleted: 'Theme deleted',
        themeMoved: 'Theme moved',
        questionAdded: 'Question added',
        questionDeleted: 'Question deleted',
        questionMoved: 'Question moved',
        fileDeleted: 'File deleted',
        uploaded: 'Uploaded',
        failedToUpload: 'Failed to upload',
        failedToSave: 'Failed to save',
        noPackageSelected: 'No package selected',

        // Confirmations
        confirmDeleteRound: 'Delete this round?',
        confirmDeleteTheme: 'Delete this theme?',
        confirmDeleteQuestion: 'Delete this question?',
        confirmDelete: 'Delete {name}?'
    },
    ru: {
        // Header
        new: 'Новый',
        open: 'Открыть',
        save: 'Сохранить',
        export: 'Экспорт',
        exportSiq: 'Экспорт в .siq',
        exportXml: 'Экспорт в XML',
        exportYaml: 'Экспорт в YAML',
        toggleTheme: 'Сменить тему',
        language: 'Язык',

        // Sidebar
        packages: 'Пакеты',
        refresh: 'Обновить',
        noPackagesOpen: 'Нет открытых пакетов',
        createNew: 'Создать новый',
        close: 'Закрыть',

        // Welcome
        welcomeTitle: 'Добро пожаловать в SIQuester',
        welcomeSubtitle: 'Кроссплатформенный редактор пакетов вопросов для SIGame',
        createNewPackage: 'Создать новый пакет',
        openExistingPackage: 'Открыть существующий пакет',

        // Tree
        packageStructure: 'Структура пакета',
        addRound: 'Добавить раунд',
        addTheme: 'Добавить тему',
        addQuestion: 'Добавить вопрос',
        delete: 'Удалить',
        media: 'Медиа',
        files: 'файлов',
        selectItemToEdit: 'Выберите элемент из дерева для редактирования',
        noPackageLoaded: 'Пакет не загружен',
        empty: '(пусто)',

        // Package Editor
        packageSettings: 'Настройки пакета',
        general: 'Общее',
        name: 'Название',
        date: 'Дата',
        packageLanguage: 'Язык',
        publisher: 'Издатель',
        difficulty: 'Сложность',
        restriction: 'Ограничение',
        restrictionPlaceholder: 'напр., 18+',
        authors: 'Авторы',
        addAuthor: '+ Добавить автора',
        comments: 'Комментарии',

        // Media Manager
        mediaManager: 'Медиа-менеджер',
        images: 'Изображения',
        audio: 'Аудио',
        video: 'Видео',
        html: 'HTML',
        uploadNewMedia: 'Загрузить медиафайл',
        noMediaFiles: 'Нет файлов типа {type}',

        // Round Editor
        round: 'Раунд',
        type: 'Тип',
        roundTypeStandard: 'Стандартный',
        roundTypeFinal: 'Финальный',

        // Theme Editor
        theme: 'Тема',

        // Question Editor
        questionEditor: 'Редактор вопроса',
        price: 'Цена',
        questionTypeDefault: 'Обычный',
        questionTypeStake: 'Ставка (Аукцион)',
        questionTypeSecret: 'Секрет (Кот в мешке)',
        questionTypeNoRisk: 'Без риска',
        questionTypeForAll: 'Для всех (Одновременный)',
        questionContent: 'Содержимое вопроса',
        text: 'Текст',
        questionText: 'Текст вопроса',
        rightAnswers: 'Правильные ответы',
        addAnswer: '+ Добавить ответ',
        wrongAnswers: 'Неправильные ответы (необязательно)',
        addWrongAnswer: '+ Добавить неправильный ответ',
        addText: '+ Текст',
        addMedia: '+ Медиа',
        mediaAdded: 'Медиа добавлено в вопрос',
        noMediaUploaded: 'Медиафайлы не загружены. Сначала загрузите медиа.',
        answerContent: 'Содержимое ответа (отображается при показе ответа)',
        answerText: 'Текст ответа',
        mediaAddedToAnswer: 'Медиа добавлено в ответ',

        // Toasts
        packageCreated: 'Пакет создан',
        packageOpened: 'Пакет открыт',
        packageClosed: 'Пакет закрыт',
        roundAdded: 'Раунд добавлен',
        roundDeleted: 'Раунд удалён',
        roundMoved: 'Раунд перемещён',
        themeAdded: 'Тема добавлена',
        themeDeleted: 'Тема удалена',
        themeMoved: 'Тема перемещена',
        questionAdded: 'Вопрос добавлен',
        questionDeleted: 'Вопрос удалён',
        questionMoved: 'Вопрос перемещён',
        fileDeleted: 'Файл удалён',
        uploaded: 'Загружено',
        failedToUpload: 'Не удалось загрузить',
        failedToSave: 'Не удалось сохранить',
        noPackageSelected: 'Пакет не выбран',

        // Confirmations
        confirmDeleteRound: 'Удалить этот раунд?',
        confirmDeleteTheme: 'Удалить эту тему?',
        confirmDeleteQuestion: 'Удалить этот вопрос?',
        confirmDelete: 'Удалить {name}?'
    }
};

// State
let state = {
    packages: [],
    currentPackageId: null,
    currentPackage: null,
    selectedNode: null,
    theme: localStorage.getItem('theme') || 'light',
    lang: localStorage.getItem('lang') || 'ru',
    media: { Images: [], Audio: [], Video: [], Html: [] }
};

// Translation function
function t(key, params = {}) {
    let text = i18n[state.lang]?.[key] || i18n['en'][key] || key;
    for (const [k, v] of Object.entries(params)) {
        text = text.replace(`{${k}}`, v);
    }
    return text;
}

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    applyTheme();
    applyLanguage();
    refreshPackages();
});

// Theme
function toggleTheme() {
    state.theme = state.theme === 'light' ? 'dark' : 'light';
    localStorage.setItem('theme', state.theme);
    applyTheme();
}

function applyTheme() {
    document.documentElement.setAttribute('data-theme', state.theme);
}

// Language
function toggleLanguage() {
    state.lang = state.lang === 'en' ? 'ru' : 'en';
    localStorage.setItem('lang', state.lang);
    applyLanguage();
    renderAll();
}

function setLanguage(lang) {
    state.lang = lang;
    localStorage.setItem('lang', state.lang);
    applyLanguage();
    renderAll();
}

function applyLanguage() {
    document.documentElement.setAttribute('lang', state.lang);
    // Update static HTML elements
    updateStaticText();
}

function updateStaticText() {
    // Header buttons
    const btnNew = document.querySelector('[data-i18n="new"]');
    const btnOpen = document.querySelector('[data-i18n="open"]');
    const btnSave = document.querySelector('[data-i18n="save"]');
    const btnExport = document.querySelector('[data-i18n="export"]');

    if (btnNew) btnNew.textContent = t('new');
    if (btnOpen) btnOpen.textContent = t('open');
    if (btnSave) btnSave.textContent = t('save');
    if (btnExport) btnExport.textContent = t('export');

    // Export dropdown
    const exportSiq = document.querySelector('[data-i18n="exportSiq"]');
    const exportXml = document.querySelector('[data-i18n="exportXml"]');
    const exportYaml = document.querySelector('[data-i18n="exportYaml"]');

    if (exportSiq) exportSiq.textContent = t('exportSiq');
    if (exportXml) exportXml.textContent = t('exportXml');
    if (exportYaml) exportYaml.textContent = t('exportYaml');

    // Sidebar
    const packagesHeader = document.querySelector('[data-i18n="packages"]');
    if (packagesHeader) packagesHeader.textContent = t('packages');

    // Welcome screen
    const welcomeTitle = document.querySelector('[data-i18n="welcomeTitle"]');
    const welcomeSubtitle = document.querySelector('[data-i18n="welcomeSubtitle"]');
    const createNewPkg = document.querySelector('[data-i18n="createNewPackage"]');
    const openExisting = document.querySelector('[data-i18n="openExistingPackage"]');

    if (welcomeTitle) welcomeTitle.textContent = t('welcomeTitle');
    if (welcomeSubtitle) welcomeSubtitle.textContent = t('welcomeSubtitle');
    if (createNewPkg) createNewPkg.textContent = t('createNewPackage');
    if (openExisting) openExisting.textContent = t('openExistingPackage');

    // Language selector
    const langBtn = document.querySelector('.lang-btn');
    if (langBtn) langBtn.textContent = state.lang.toUpperCase();
}

function renderAll() {
    updateStaticText();
    renderPackageList();
    if (state.currentPackage) {
        renderTreeView();
        renderEditor();
    }
}

// Toast notifications
function showToast(message, type = 'success') {
    const container = document.getElementById('toast-container');
    const toast = document.createElement('div');
    toast.className = `toast toast-${type}`;
    toast.textContent = message;
    container.appendChild(toast);
    setTimeout(() => {
        toast.remove();
    }, 3000);
}

// API helpers
async function api(endpoint, options = {}) {
    const url = `${API_BASE}${endpoint}`;
    const response = await fetch(url, {
        headers: {
            'Content-Type': 'application/json',
            ...options.headers
        },
        ...options
    });
    if (!response.ok) {
        const error = await response.json().catch(() => ({ error: 'Request failed' }));
        throw new Error(error.error || 'Request failed');
    }
    return response.json();
}

// Package operations
async function refreshPackages() {
    try {
        state.packages = await api('/packages');
        renderPackageList();
    } catch (err) {
        console.error('Failed to refresh packages:', err);
    }
}

async function createPackage() {
    try {
        const result = await api('/packages', {
            method: 'POST',
            body: JSON.stringify({ name: 'New Package', author: 'Author' })
        });
        state.packages.push({ id: result.id, name: result.package.name });
        renderPackageList();
        selectPackage(result.id);
        showToast(t('packageCreated'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function selectPackage(id) {
    try {
        const result = await api(`/packages/${id}`);
        state.currentPackageId = id;
        state.currentPackage = result.package;
        state.selectedNode = null;
        // Load media lists
        await loadMediaLists();
        // Update UI
        renderPackageList();
        showEditorArea();
        renderTreeView();
        renderEditor();
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function loadMediaLists() {
    if (!state.currentPackageId) return;
    try {
        const types = ['Images', 'Audio', 'Video', 'Html'];
        for (const type of types) {
            const result = await api(`/packages/${state.currentPackageId}/media/${type}`);
            state.media[type] = result.files || [];
        }
    } catch (err) {
        console.error('Failed to load media:', err);
    }
}

async function closePackage(id, event) {
    event.stopPropagation();
    try {
        await api(`/packages/${id}`, { method: 'DELETE' });
        state.packages = state.packages.filter(p => p.id !== id);
        if (state.currentPackageId === id) {
            state.currentPackageId = null;
            state.currentPackage = null;
            state.selectedNode = null;
            state.media = { Images: [], Audio: [], Video: [], Html: [] };
            hideEditorArea();
        }
        renderPackageList();
        showToast(t('packageClosed'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function saveCurrentPackage() {
    if (!state.currentPackageId) {
        showToast(t('noPackageSelected'), 'warning');
        return;
    }
    // For web version, we download the file
    exportSIQ();
}

// File operations
function openFileDialog() {
    document.getElementById('file-input').click();
}

async function handleFileSelect(event) {
    const file = event.target.files[0];
    if (!file) return;
    const formData = new FormData();
    formData.append('file', file);
    try {
        const response = await fetch(`${API_BASE}/packages/open`, {
            method: 'POST',
            body: formData
        });
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Failed to open file');
        }
        const result = await response.json();
        state.packages.push({ id: result.id, name: result.package.name, filePath: file.name });
        renderPackageList();
        selectPackage(result.id);
        showToast(t('packageOpened'));
    } catch (err) {
        showToast(err.message, 'error');
    }
    event.target.value = '';
}

// Export
function exportSIQ() {
    if (!state.currentPackageId) return;
    window.location.href = `${API_BASE}/export/${state.currentPackageId}/siq`;
}

function exportXML() {
    if (!state.currentPackageId) return;
    window.location.href = `${API_BASE}/export/${state.currentPackageId}/xml`;
}

function exportYAML() {
    if (!state.currentPackageId) return;
    window.location.href = `${API_BASE}/export/${state.currentPackageId}/yaml`;
}

// UI Rendering
function renderPackageList() {
    const list = document.getElementById('package-list');
    if (state.packages.length === 0) {
        list.innerHTML = `
            <div class="empty-state">
                <p>${t('noPackagesOpen')}</p>
                <button class="btn btn-primary btn-small" onclick="createPackage()">${t('createNew')}</button>
            </div>
        `;
        return;
    }
    list.innerHTML = state.packages.map(pkg => `
        <div class="package-item ${pkg.id === state.currentPackageId ? 'active' : ''}"
             onclick="selectPackage('${pkg.id}')">
            <div class="package-item-info">
                <span class="package-item-name">${escapeHtml(pkg.name)}</span>
                ${pkg.filePath ? `<span class="package-item-path">${escapeHtml(pkg.filePath)}</span>` : ''}
            </div>
            <button class="package-item-close" onclick="closePackage('${pkg.id}', event)" title="${t('close')}">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"/>
                    <line x1="6" y1="6" x2="18" y2="18"/>
                </svg>
            </button>
        </div>
    `).join('');
}

function showEditorArea() {
    document.getElementById('welcome-screen').classList.add('hidden');
    document.getElementById('editor-area').classList.remove('hidden');
}

function hideEditorArea() {
    document.getElementById('welcome-screen').classList.remove('hidden');
    document.getElementById('editor-area').classList.add('hidden');
}

function renderTreeView() {
    const treeView = document.getElementById('tree-view');
    const pkg = state.currentPackage;
    if (!pkg) {
        treeView.innerHTML = `<div class="empty-state"><p>${t('noPackageLoaded')}</p></div>`;
        return;
    }
    document.getElementById('tree-title').textContent = pkg.name;
    let html = '';
    // Package node
    html += `
        <div class="tree-node">
            <div class="tree-item ${state.selectedNode?.type === 'package' ? 'selected' : ''}" data-type="package">
                <span class="tree-toggle expanded">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="9 18 15 12 9 6"/>
                    </svg>
                </span>
                <span class="tree-item-icon">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"/>
                    </svg>
                </span>
                <span class="tree-item-label">${escapeHtml(pkg.name)}</span>
            </div>
            <div class="tree-children">
    `;
    // Media node
    const totalMedia = state.media.Images.length + state.media.Audio.length + state.media.Video.length + state.media.Html.length;
    html += `
        <div class="tree-node">
            <div class="tree-item ${state.selectedNode?.type === 'media' ? 'selected' : ''}" data-type="media">
                <span class="tree-toggle expanded">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="9 18 15 12 9 6"/>
                    </svg>
                </span>
                <span class="tree-item-icon">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                        <circle cx="8.5" cy="8.5" r="1.5"/>
                        <polyline points="21 15 16 10 5 21"/>
                    </svg>
                </span>
                <span class="tree-item-label">${t('media')} (${totalMedia} ${t('files')})</span>
            </div>
        </div>
    `;
    // Rounds
    if (pkg.rounds && pkg.rounds.length > 0) {
        for (let ri = 0; ri < pkg.rounds.length; ri++) {
            const round = pkg.rounds[ri];
            const roundSelected = state.selectedNode?.type === 'round' && state.selectedNode?.roundIndex === ri;
            html += `
                <div class="tree-node" data-round-node="${ri}">
                    <div class="tree-item draggable ${roundSelected ? 'selected' : ''}" data-type="round" data-round="${ri}"
                         draggable="true" data-drag-type="round" data-drag-index="${ri}">
                        <span class="tree-toggle expanded">
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <polyline points="9 18 15 12 9 6"/>
                            </svg>
                        </span>
                        <span class="tree-item-icon drag-handle" title="Drag to reorder">
                            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <circle cx="12" cy="12" r="10"/>
                            </svg>
                        </span>
                        <span class="tree-item-label">${escapeHtml(round.name)}</span>
                        <div class="tree-item-actions">
                            <button class="btn btn-icon btn-small" data-action="add-theme" data-round="${ri}" title="${t('addTheme')}">
                                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <line x1="12" y1="5" x2="12" y2="19"/>
                                    <line x1="5" y1="12" x2="19" y2="12"/>
                                </svg>
                            </button>
                            <button class="btn btn-icon btn-small" data-action="delete-round" data-round="${ri}" title="${t('delete')}">
                                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <polyline points="3 6 5 6 21 6"/>
                                    <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                                </svg>
                            </button>
                        </div>
                    </div>
                    <div class="tree-children">
            `;
            // Themes
            if (round.themes && round.themes.length > 0) {
                for (let ti = 0; ti < round.themes.length; ti++) {
                    const theme = round.themes[ti];
                    const themeSelected = state.selectedNode?.type === 'theme' &&
                        state.selectedNode?.roundIndex === ri && state.selectedNode?.themeIndex === ti;
                    html += `
                        <div class="tree-node" data-theme-node="${ti}" data-parent-round="${ri}">
                            <div class="tree-item draggable ${themeSelected ? 'selected' : ''}" data-type="theme" data-round="${ri}" data-theme="${ti}"
                                 draggable="true" data-drag-type="theme" data-drag-round="${ri}" data-drag-theme="${ti}">
                                <span class="tree-toggle expanded">
                                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                        <polyline points="9 18 15 12 9 6"/>
                                    </svg>
                                </span>
                                <span class="tree-item-icon drag-handle" title="Drag to move">
                                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                        <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                                        <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
                                    </svg>
                                </span>
                                <span class="tree-item-label">${escapeHtml(theme.name)}</span>
                                <div class="tree-item-actions">
                                    <button class="btn btn-icon btn-small" data-action="add-question" data-round="${ri}" data-theme="${ti}" title="${t('addQuestion')}">
                                        <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                            <line x1="12" y1="5" x2="12" y2="19"/>
                                            <line x1="5" y1="12" x2="19" y2="12"/>
                                        </svg>
                                    </button>
                                    <button class="btn btn-icon btn-small" data-action="delete-theme" data-round="${ri}" data-theme="${ti}" title="${t('delete')}">
                                        <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                            <polyline points="3 6 5 6 21 6"/>
                                            <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                            <div class="tree-children">
                    `;
                    // Questions
                    if (theme.questions && theme.questions.length > 0) {
                        for (let qi = 0; qi < theme.questions.length; qi++) {
                            const question = theme.questions[qi];
                            const qSelected = state.selectedNode?.type === 'question' &&
                                state.selectedNode?.roundIndex === ri &&
                                state.selectedNode?.themeIndex === ti &&
                                state.selectedNode?.questionIndex === qi;
                            html += `
                                <div class="tree-node" data-question-node="${qi}" data-parent-round="${ri}" data-parent-theme="${ti}">
                                    <div class="tree-item draggable ${qSelected ? 'selected' : ''}" data-type="question" data-round="${ri}" data-theme="${ti}" data-question="${qi}"
                                         draggable="true" data-drag-type="question" data-drag-round="${ri}" data-drag-theme="${ti}" data-drag-question="${qi}">
                                        <span class="tree-toggle" style="visibility: hidden;">
                                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                                <polyline points="9 18 15 12 9 6"/>
                                            </svg>
                                        </span>
                                        <span class="tree-item-icon drag-handle" title="Drag to reorder">
                                            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                                <circle cx="12" cy="12" r="10"/>
                                                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/>
                                                <line x1="12" y1="17" x2="12.01" y2="17"/>
                                            </svg>
                                        </span>
                                        <span class="tree-item-label">${question.price}: ${escapeHtml(getQuestionPreview(question))}</span>
                                        <div class="tree-item-actions">
                                            <button class="btn btn-icon btn-small" data-action="delete-question" data-round="${ri}" data-theme="${ti}" data-question="${qi}" title="${t('delete')}">
                                                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                                    <polyline points="3 6 5 6 21 6"/>
                                                    <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                                                </svg>
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            `;
                        }
                    }
                    html += `
                            </div>
                        </div>
                    `;
                }
            }
            html += `
                    </div>
                </div>
            `;
        }
    }
    html += `
            </div>
        </div>
    `;
    treeView.innerHTML = html;
    // Add event listeners using event delegation (only once)
    if (!treeView.dataset.listenersInitialized) {
        treeView.addEventListener('click', handleTreeClick);
        // Initialize drag and drop
        initTreeDragDrop(treeView);
        treeView.dataset.listenersInitialized = 'true';
    }
}

function handleTreeClick(event) {
    const target = event.target;
    // Handle toggle clicks
    const toggle = target.closest('.tree-toggle');
    if (toggle && toggle.style.visibility !== 'hidden') {
        event.stopPropagation();
        toggle.classList.toggle('expanded');
        const children = toggle.closest('.tree-item').nextElementSibling;
        if (children && children.classList.contains('tree-children')) {
            children.style.display = toggle.classList.contains('expanded') ? 'block' : 'none';
        }
        return;
    }
    // Handle action button clicks
    const actionBtn = target.closest('[data-action]');
    if (actionBtn) {
        event.stopPropagation();
        const action = actionBtn.dataset.action;
        const ri = parseInt(actionBtn.dataset.round);
        const ti = parseInt(actionBtn.dataset.theme);
        const qi = parseInt(actionBtn.dataset.question);
        switch (action) {
            case 'add-theme': addTheme(ri); break;
            case 'delete-round': deleteRound(ri); break;
            case 'add-question': addQuestion(ri, ti); break;
            case 'delete-theme': deleteTheme(ri, ti); break;
            case 'delete-question': deleteQuestion(ri, ti, qi); break;
        }
        return;
    }
    // Handle tree item selection
    const treeItem = target.closest('.tree-item');
    if (treeItem) {
        const type = treeItem.dataset.type;
        const ri = parseInt(treeItem.dataset.round);
        const ti = parseInt(treeItem.dataset.theme);
        const qi = parseInt(treeItem.dataset.question);
        selectNode(type, ri, ti, qi);
    }
}

function selectNode(type, roundIndex, themeIndex, questionIndex) {
    state.selectedNode = { type, roundIndex, themeIndex, questionIndex };
    renderTreeView();
    renderEditor();
}

// Drag and Drop
let dragState = {
    dragging: null,
    dragType: null,
    dragRound: null,
    dragTheme: null,
    dragQuestion: null
};

function initTreeDragDrop(treeView) {
    // Drag start
    treeView.addEventListener('dragstart', (e) => {
        const item = e.target.closest('.tree-item[draggable="true"]');
        if (!item) return;
        dragState.dragging = item;
        dragState.dragType = item.dataset.dragType;
        dragState.dragRound = parseInt(item.dataset.dragRound || item.dataset.dragIndex);
        dragState.dragTheme = parseInt(item.dataset.dragTheme);
        dragState.dragQuestion = parseInt(item.dataset.dragQuestion);
        item.classList.add('dragging');
        e.dataTransfer.effectAllowed = 'move';
        e.dataTransfer.setData('text/plain', JSON.stringify({
            type: dragState.dragType,
            round: dragState.dragRound,
            theme: dragState.dragTheme,
            question: dragState.dragQuestion
        }));
    });
    // Drag end
    treeView.addEventListener('dragend', (e) => {
        const item = e.target.closest('.tree-item');
        if (item) item.classList.remove('dragging');
        clearDropIndicators();
        dragState = { dragging: null, dragType: null, dragRound: null, dragTheme: null, dragQuestion: null };
    });
    // Drag over
    treeView.addEventListener('dragover', (e) => {
        e.preventDefault();
        const item = e.target.closest('.tree-item');
        if (!item || !dragState.dragging || item === dragState.dragging) return;
        e.dataTransfer.dropEffect = 'move';
        const targetType = item.dataset.type;
        const rect = item.getBoundingClientRect();
        const midY = rect.top + rect.height / 2;
        clearDropIndicators();
        // Determine drop position
        if (dragState.dragType === 'round' && targetType === 'round') {
            // Reordering rounds - show above/below indicator
            if (e.clientY < midY) {
                item.classList.add('drag-over-above');
            } else {
                item.classList.add('drag-over-below');
            }
        } else if (dragState.dragType === 'theme') {
            if (targetType === 'theme') {
                // Moving theme within/between rounds - show above/below indicator
                if (e.clientY < midY) {
                    item.classList.add('drag-over-above');
                } else {
                    item.classList.add('drag-over-below');
                }
            } else if (targetType === 'round') {
                // Dropping theme into a round
                item.classList.add('drag-over');
            }
        } else if (dragState.dragType === 'question') {
            if (targetType === 'question') {
                // Reordering questions within same theme
                const targetRound = parseInt(item.dataset.round);
                const targetTheme = parseInt(item.dataset.theme);
                // Only allow reordering within the same theme
                if (dragState.dragRound === targetRound && dragState.dragTheme === targetTheme) {
                    if (e.clientY < midY) {
                        item.classList.add('drag-over-above');
                    } else {
                        item.classList.add('drag-over-below');
                    }
                }
            }
        }
    });
    // Drag leave
    treeView.addEventListener('dragleave', (e) => {
        const item = e.target.closest('.tree-item');
        if (item) {
            item.classList.remove('drag-over', 'drag-over-above', 'drag-over-below');
        }
    });
    // Drop
    treeView.addEventListener('drop', async (e) => {
        e.preventDefault();
        e.stopPropagation();
        const item = e.target.closest('.tree-item');
        if (!item || !dragState.dragging || item === dragState.dragging) {
            clearDropIndicators();
            return;
        }
        const targetType = item.dataset.type;
        const targetRound = parseInt(item.dataset.round);
        const targetTheme = parseInt(item.dataset.theme);
        const targetQuestion = parseInt(item.dataset.question);
        const rect = item.getBoundingClientRect();
        const midY = rect.top + rect.height / 2;
        const insertBefore = e.clientY < midY;
        clearDropIndicators();
        try {
            if (dragState.dragType === 'round' && targetType === 'round') {
                // Reorder rounds
                const fromIndex = dragState.dragRound;
                let toIndex = targetRound;
                if (!insertBefore) toIndex++;
                if (fromIndex < toIndex) toIndex--;
                if (fromIndex !== toIndex) {
                    await reorderRound(fromIndex, toIndex);
                }
            } else if (dragState.dragType === 'theme') {
                if (targetType === 'theme') {
                    // Move theme to new position
                    const fromRound = dragState.dragRound;
                    const fromTheme = dragState.dragTheme;
                    let toRound = targetRound;
                    let toTheme = targetTheme;
                    if (!insertBefore) toTheme++;
                    if (fromRound === toRound && fromTheme < toTheme) toTheme--;
                    if (fromRound !== toRound || fromTheme !== toTheme) {
                        await moveTheme(fromRound, fromTheme, toRound, toTheme);
                    }
                } else if (targetType === 'round') {
                    // Move theme into round (at the end)
                    const fromRound = dragState.dragRound;
                    const fromTheme = dragState.dragTheme;
                    const toRound = targetRound;
                    if (fromRound !== toRound) {
                        const toTheme = state.currentPackage.rounds[toRound].themes?.length || 0;
                        await moveTheme(fromRound, fromTheme, toRound, toTheme);
                    }
                }
            } else if (dragState.dragType === 'question' && targetType === 'question') {
                // Reorder questions within the same theme
                if (dragState.dragRound === targetRound && dragState.dragTheme === targetTheme) {
                    const fromIndex = dragState.dragQuestion;
                    let toIndex = targetQuestion;
                    if (!insertBefore) toIndex++;
                    if (fromIndex < toIndex) toIndex--;
                    if (fromIndex !== toIndex) {
                        await reorderQuestion(targetRound, targetTheme, fromIndex, toIndex);
                    }
                }
            }
        } catch (err) {
            showToast(err.message, 'error');
        }
    });
}

function clearDropIndicators() {
    document.querySelectorAll('.drag-over, .drag-over-above, .drag-over-below').forEach(el => {
        el.classList.remove('drag-over', 'drag-over-above', 'drag-over-below');
    });
}

async function reorderRound(fromIndex, toIndex) {
    if (!state.currentPackageId) return;
    try {
        await api(`/packages/${state.currentPackageId}/rounds/reorder`, {
            method: 'POST',
            body: JSON.stringify({ from: fromIndex, to: toIndex })
        });
        // Update local state
        const rounds = state.currentPackage.rounds;
        const [moved] = rounds.splice(fromIndex, 1);
        rounds.splice(toIndex, 0, moved);
        // Update selected node if needed
        if (state.selectedNode?.type === 'round') {
            if (state.selectedNode.roundIndex === fromIndex) {
                state.selectedNode.roundIndex = toIndex;
            } else if (fromIndex < state.selectedNode.roundIndex && state.selectedNode.roundIndex <= toIndex) {
                state.selectedNode.roundIndex--;
            } else if (toIndex <= state.selectedNode.roundIndex && state.selectedNode.roundIndex < fromIndex) {
                state.selectedNode.roundIndex++;
            }
        }
        renderTreeView();
        showToast(t('roundMoved') || 'Round moved');
    } catch (err) {
        throw err;
    }
}

async function moveTheme(fromRound, fromTheme, toRound, toTheme) {
    if (!state.currentPackageId) return;
    try {
        await api(`/packages/${state.currentPackageId}/themes/move`, {
            method: 'POST',
            body: JSON.stringify({ fromRound, fromTheme, toRound, toTheme })
        });
        // Update local state
        const theme = state.currentPackage.rounds[fromRound].themes.splice(fromTheme, 1)[0];
        if (!state.currentPackage.rounds[toRound].themes) {
            state.currentPackage.rounds[toRound].themes = [];
        }
        state.currentPackage.rounds[toRound].themes.splice(toTheme, 0, theme);
        // Update selected node if needed
        if (state.selectedNode?.type === 'theme' || state.selectedNode?.type === 'question') {
            if (state.selectedNode.roundIndex === fromRound && state.selectedNode.themeIndex === fromTheme) {
                state.selectedNode.roundIndex = toRound;
                state.selectedNode.themeIndex = toTheme;
            }
        }
        renderTreeView();
        showToast(t('themeMoved') || 'Theme moved');
    } catch (err) {
        throw err;
    }
}
async function reorderQuestion(roundIndex, themeIndex, fromIndex, toIndex) {
    if (!state.currentPackageId) return;
    try {
        await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}/themes/${themeIndex}/questions/reorder`, {
            method: 'POST',
            body: JSON.stringify({ from: fromIndex, to: toIndex })
        });
        // Update local state
        const questions = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions;
        const [moved] = questions.splice(fromIndex, 1);
        questions.splice(toIndex, 0, moved);
        // Update selected node if needed
        if (state.selectedNode?.type === 'question' &&
            state.selectedNode.roundIndex === roundIndex &&
            state.selectedNode.themeIndex === themeIndex) {
            if (state.selectedNode.questionIndex === fromIndex) {
                state.selectedNode.questionIndex = toIndex;
            } else if (fromIndex < state.selectedNode.questionIndex && state.selectedNode.questionIndex <= toIndex) {
                state.selectedNode.questionIndex--;
            } else if (toIndex <= state.selectedNode.questionIndex && state.selectedNode.questionIndex < fromIndex) {
                state.selectedNode.questionIndex++;
            }
        }
        renderTreeView();
        showToast(t('questionMoved'));
    } catch (err) {
        throw err;
    }
}

function getQuestionPreview(question) {
    if (question.parameters?.question?.contentValue) {
        for (const item of question.parameters.question.contentValue) {
            if (item.type === 'text' || !item.type) {
                return item.value.substring(0, 50) + (item.value.length > 50 ? '...' : '');
            }
        }
    }
    // Also handle simpleValue (from XML parsing)
    if (question.parameters?.question?.simpleValue) {
        const text = question.parameters.question.simpleValue.trim();
        return text.substring(0, 50) + (text.length > 50 ? '...' : '');
    }
    if (question.right && question.right.length > 0) {
        return `(${question.right[0]})`;
    }
    return t('empty');
}

// Editor rendering
function renderEditor() {
    const editor = document.getElementById('editor-content');
    const node = state.selectedNode;
    const pkg = state.currentPackage;
    if (!node || !pkg) {
        editor.innerHTML = `<div class="empty-state"><p>${t('selectItemToEdit')}</p></div>`;
        return;
    }
    switch (node.type) {
        case 'package':
            renderPackageEditor(editor, pkg);
            break;
        case 'media':
            renderMediaManager(editor);
            break;
        case 'round':
            if (node.roundIndex >= 0 && node.roundIndex < pkg.rounds.length) {
                renderRoundEditor(editor, pkg.rounds[node.roundIndex], node.roundIndex);
            }
            break;
        case 'theme':
            if (node.roundIndex >= 0 && node.roundIndex < pkg.rounds.length &&
                node.themeIndex >= 0 && node.themeIndex < pkg.rounds[node.roundIndex].themes.length) {
                renderThemeEditor(editor, pkg.rounds[node.roundIndex].themes[node.themeIndex], node.roundIndex, node.themeIndex);
            }
            break;
        case 'question':
            if (node.roundIndex >= 0 && node.roundIndex < pkg.rounds.length &&
                node.themeIndex >= 0 && node.themeIndex < pkg.rounds[node.roundIndex].themes.length &&
                node.questionIndex >= 0 && node.questionIndex < pkg.rounds[node.roundIndex].themes[node.themeIndex].questions.length) {
                renderQuestionEditor(editor,
                    pkg.rounds[node.roundIndex].themes[node.themeIndex].questions[node.questionIndex],
                    node.roundIndex, node.themeIndex, node.questionIndex);
            }
            break;
    }
}

function renderPackageEditor(editor, pkg) {
    editor.innerHTML = `
        <h2 style="margin-bottom: 24px;">${t('packageSettings')}</h2>
        <div class="form-section">
            <div class="form-section-title">${t('general')}</div>
            <div class="form-group">
                <label class="form-label">${t('name')}</label>
                <input type="text" class="form-input" value="${escapeHtml(pkg.name)}"
                       onchange="updatePackageField('name', this.value)">
            </div>
            <div class="form-row">
                <div class="form-group">
                    <label class="form-label">${t('date')}</label>
                    <input type="text" class="form-input" value="${escapeHtml(pkg.date || '')}"
                           onchange="updatePackageField('date', this.value)">
                </div>
                <div class="form-group">
                    <label class="form-label">${t('packageLanguage')}</label>
                    <input type="text" class="form-input" value="${escapeHtml(pkg.language || '')}"
                           onchange="updatePackageField('language', this.value)">
                </div>
            </div>
            <div class="form-row">
                <div class="form-group">
                    <label class="form-label">${t('publisher')}</label>
                    <input type="text" class="form-input" value="${escapeHtml(pkg.publisher || '')}"
                           onchange="updatePackageField('publisher', this.value)">
                </div>
                <div class="form-group">
                    <label class="form-label">${t('difficulty')} (0-10)</label>
                    <input type="number" class="form-input" min="0" max="10" value="${pkg.difficulty || 5}"
                           onchange="updatePackageField('difficulty', parseInt(this.value))">
                </div>
            </div>
            <div class="form-group">
                <label class="form-label">${t('restriction')}</label>
                <input type="text" class="form-input" value="${escapeHtml(pkg.restriction || '')}"
                       placeholder="${t('restrictionPlaceholder')}"
                       onchange="updatePackageField('restriction', this.value)">
            </div>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('authors')}</div>
            <div id="authors-list" class="answer-list">
                ${(pkg.info?.authors || []).map((author, i) => `
                    <div class="answer-item">
                        <input type="text" class="form-input" value="${escapeHtml(author)}"
                               onchange="updatePackageAuthor(${i}, this.value)">
                        <button class="btn btn-icon" onclick="removePackageAuthor(${i})">
                            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                `).join('')}
            </div>
            <button class="btn btn-secondary btn-small" style="margin-top: 8px;" onclick="addPackageAuthor()">${t('addAuthor')}</button>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('comments')}</div>
            <div class="form-group">
                <textarea class="form-input" rows="3" onchange="updatePackageField('info.comments', this.value)">${escapeHtml(pkg.info?.comments || '')}</textarea>
            </div>
        </div>
    `;
}

function renderMediaManager(editor) {
    const types = [
        { key: 'Images', label: t('images'), icon: 'image' },
        { key: 'Audio', label: t('audio'), icon: 'music' },
        { key: 'Video', label: t('video'), icon: 'video' },
        { key: 'Html', label: t('html'), icon: 'code' }
    ];
    editor.innerHTML = `
        <h2 style="margin-bottom: 24px;">${t('mediaManager')}</h2>
        <div class="tabs">
            ${types.map((type, i) => `
                <button class="tab ${i === 0 ? 'active' : ''}" onclick="switchMediaTab('${type.key}', this)">${type.label} (${state.media[type.key].length})</button>
            `).join('')}
        </div>
        <div class="form-section">
            <div class="form-group">
                <label class="form-label">${t('uploadNewMedia')}</label>
                <input type="file" id="media-upload" class="form-input" onchange="uploadMediaFile(event)" multiple>
            </div>
        </div>
        <div id="media-content">
            ${renderMediaGrid('Images')}
        </div>
    `;
}

function switchMediaTab(type, btn) {
    document.querySelectorAll('.tabs .tab').forEach(t => t.classList.remove('active'));
    btn.classList.add('active');
    document.getElementById('media-content').innerHTML = renderMediaGrid(type);
}

function renderMediaGrid(type) {
    const files = state.media[type] || [];
    if (files.length === 0) {
        const typeLabel = type === 'Images' ? t('images').toLowerCase() :
                         type === 'Audio' ? t('audio').toLowerCase() :
                         type === 'Video' ? t('video').toLowerCase() : 'HTML';
        return `<div class="empty-state"><p>${t('noMediaFiles', { type: typeLabel })}</p></div>`;
    }
    return `
        <div class="media-grid">
            ${files.map(name => {
                const url = `${API_BASE}/packages/${state.currentPackageId}/media/${type}/${encodeURIComponent(name)}`;
                let preview = '';
                if (type === 'Images') {
                    preview = `<img src="${url}" alt="${escapeHtml(name)}" loading="lazy">`;
                } else if (type === 'Audio') {
                    preview = `<div class="media-placeholder">
                        <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M9 18V5l12-2v13"/>
                            <circle cx="6" cy="18" r="3"/>
                            <circle cx="18" cy="16" r="3"/>
                        </svg>
                        <audio controls src="${url}" style="width:100%;margin-top:8px;"></audio>
                    </div>`;
                } else if (type === 'Video') {
                    preview = `<video src="${url}" controls style="width:100%;height:100%;object-fit:cover;"></video>`;
                } else {
                    preview = `<div class="media-placeholder">
                        <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <polyline points="16 18 22 12 16 6"/>
                            <polyline points="8 6 2 12 8 18"/>
                        </svg>
                    </div>`;
                }
                return `
                    <div class="media-item" onclick="previewMedia('${type}', '${escapeHtml(name)}')">
                        ${preview}
                        <span class="media-item-name">${escapeHtml(name)}</span>
                        <div class="media-item-overlay">
                            <button class="btn btn-danger btn-small" onclick="deleteMediaFile('${type}', '${escapeHtml(name)}'); event.stopPropagation();">${t('delete')}</button>
                        </div>
                    </div>
                `;
            }).join('')}
        </div>
    `;
}

async function uploadMediaFile(event) {
    const files = event.target.files;
    if (!files || files.length === 0) return;
    // Determine type from active tab
    const activeTab = document.querySelector('.tabs .tab.active');
    const type = activeTab ? activeTab.textContent.split(' ')[0] : 'Images';
    // Map translated label back to key
    const typeMap = {
        [t('images')]: 'Images',
        [t('audio')]: 'Audio',
        [t('video')]: 'Video',
        [t('html')]: 'Html'
    };
    const actualType = typeMap[type] || 'Images';
    for (const file of files) {
        const formData = new FormData();
        formData.append('file', file);
        try {
            await fetch(`${API_BASE}/packages/${state.currentPackageId}/media/${actualType}`, {
                method: 'POST',
                body: formData
            });
            showToast(`${t('uploaded')} ${file.name}`);
        } catch (err) {
            showToast(`${t('failedToUpload')} ${file.name}`, 'error');
        }
    }
    await loadMediaLists();
    document.getElementById('media-content').innerHTML = renderMediaGrid(actualType);
    event.target.value = '';
}

async function deleteMediaFile(type, name) {
    if (!confirm(t('confirmDelete', { name }))) return;
    try {
        await api(`/packages/${state.currentPackageId}/media/${type}/${encodeURIComponent(name)}`, {
            method: 'DELETE'
        });
        await loadMediaLists();
        document.getElementById('media-content').innerHTML = renderMediaGrid(type);
        showToast(t('fileDeleted'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

function previewMedia(type, name) {
    const url = `${API_BASE}/packages/${state.currentPackageId}/media/${type}/${encodeURIComponent(name)}`;
    window.open(url, '_blank');
}

function renderRoundEditor(editor, round, roundIndex) {
    editor.innerHTML = `
        <h2 style="margin-bottom: 24px;">${t('round')}: ${escapeHtml(round.name)}</h2>
        <div class="form-section">
            <div class="form-group">
                <label class="form-label">${t('name')}</label>
                <input type="text" class="form-input" value="${escapeHtml(round.name)}"
                       onchange="updateRound(${roundIndex}, 'name', this.value)">
            </div>
            <div class="form-group">
                <label class="form-label">${t('type')}</label>
                <select class="form-input" onchange="updateRound(${roundIndex}, 'type', this.value)">
                    <option value="standart" ${round.type === 'standart' ? 'selected' : ''}>${t('roundTypeStandard')}</option>
                    <option value="final" ${round.type === 'final' ? 'selected' : ''}>${t('roundTypeFinal')}</option>
                </select>
            </div>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('comments')}</div>
            <div class="form-group">
                <textarea class="form-input" rows="3" onchange="updateRoundInfo(${roundIndex}, 'comments', this.value)">${escapeHtml(round.info?.comments || '')}</textarea>
            </div>
        </div>
    `;
}

function renderThemeEditor(editor, theme, roundIndex, themeIndex) {
    editor.innerHTML = `
        <h2 style="margin-bottom: 24px;">${t('theme')}: ${escapeHtml(theme.name)}</h2>
        <div class="form-section">
            <div class="form-group">
                <label class="form-label">${t('name')}</label>
                <input type="text" class="form-input" value="${escapeHtml(theme.name)}"
                       onchange="updateTheme(${roundIndex}, ${themeIndex}, 'name', this.value)">
            </div>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('comments')}</div>
            <div class="form-group">
                <textarea class="form-input" rows="3" onchange="updateThemeInfo(${roundIndex}, ${themeIndex}, 'comments', this.value)">${escapeHtml(theme.info?.comments || '')}</textarea>
            </div>
        </div>
    `;
}

function renderQuestionEditor(editor, question, roundIndex, themeIndex, questionIndex) {
    const questionText = getQuestionText(question);
    editor.innerHTML = `
        <h2 style="margin-bottom: 24px;">${t('questionEditor')}</h2>
        <div class="form-section">
            <div class="form-row">
                <div class="form-group">
                    <label class="form-label">${t('price')}</label>
                    <input type="number" class="form-input" value="${question.price}" step="100"
                           onchange="updateQuestion(${roundIndex}, ${themeIndex}, ${questionIndex}, 'price', parseInt(this.value))">
                </div>
                <div class="form-group">
                    <label class="form-label">${t('type')}</label>
                    <select class="form-input" onchange="updateQuestion(${roundIndex}, ${themeIndex}, ${questionIndex}, 'typeName', this.value)">
                        <option value="default" ${question.typeName === 'default' ? 'selected' : ''}>${t('questionTypeDefault')}</option>
                        <option value="stake" ${question.typeName === 'stake' ? 'selected' : ''}>${t('questionTypeStake')}</option>
                        <option value="secret" ${question.typeName === 'secret' ? 'selected' : ''}>${t('questionTypeSecret')}</option>
                        <option value="noRisk" ${question.typeName === 'noRisk' ? 'selected' : ''}>${t('questionTypeNoRisk')}</option>
                        <option value="forAll" ${question.typeName === 'forAll' ? 'selected' : ''}>${t('questionTypeForAll')}</option>
                    </select>
                </div>
            </div>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('questionContent')}</div>
            ${renderQuestionContent(question, roundIndex, themeIndex, questionIndex)}
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('answerContent')}</div>
            ${renderAnswerContent(question, roundIndex, themeIndex, questionIndex)}
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('rightAnswers')}</div>
            <div id="right-answers" class="answer-list">
                ${(question.right || ['']).map((answer, i) => `
                    <div class="answer-item">
                        <input type="text" class="form-input" value="${escapeHtml(answer)}"
                               onchange="updateQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'right', ${i}, this.value)">
                        <button class="btn btn-icon" onclick="removeQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'right', ${i})">
                            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                `).join('')}
            </div>
            <button class="btn btn-secondary btn-small" style="margin-top: 8px;"
                    onclick="addQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'right')">${t('addAnswer')}</button>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('wrongAnswers')}</div>
            <div id="wrong-answers" class="answer-list">
                ${(question.wrong || []).map((answer, i) => `
                    <div class="answer-item">
                        <input type="text" class="form-input" value="${escapeHtml(answer)}"
                               onchange="updateQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'wrong', ${i}, this.value)">
                        <button class="btn btn-icon" onclick="removeQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'wrong', ${i})">
                            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                `).join('')}
            </div>
            <button class="btn btn-secondary btn-small" style="margin-top: 8px;"
                    onclick="addQuestionAnswer(${roundIndex}, ${themeIndex}, ${questionIndex}, 'wrong')">${t('addWrongAnswer')}</button>
        </div>
        <div class="form-section">
            <div class="form-section-title">${t('comments')}</div>
            <div class="form-group">
                <textarea class="form-input" rows="2"
                          onchange="updateQuestionInfo(${roundIndex}, ${themeIndex}, ${questionIndex}, 'comments', this.value)">${escapeHtml(question.info?.comments || '')}</textarea>
            </div>
        </div>
    `;
}

function renderQuestionContent(question, ri, ti, qi) {
    // Handle both contentValue array and simpleValue string
    let content = question.parameters?.question?.contentValue || [];
    // If no contentValue but has simpleValue, convert it to content item
    if (content.length === 0 && question.parameters?.question?.simpleValue) {
        content = [{ type: 'text', value: question.parameters.question.simpleValue.trim() }];
    }
    let html = '<div class="question-content-list">';
    for (let i = 0; i < content.length; i++) {
        const item = content[i];
        const type = item.type || 'text';
        if (type === 'text') {
            html += `
                <div class="content-item">
                    <div class="content-item-header">
                        <span class="content-item-type">${t('text')}</span>
                        <button class="btn btn-icon btn-small" onclick="removeContentItem(${ri}, ${ti}, ${qi}, ${i})" title="${t('delete')}">
                            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    <textarea class="form-input" rows="3" onchange="updateQuestionContentItem(${ri}, ${ti}, ${qi}, ${i}, 'value', this.value)">${escapeHtml(item.value || '')}</textarea>
                </div>
            `;
        } else if (item.isRef) {
            // Media reference
            const mediaType = type === 'image' ? 'Images' : type === 'audio' ? 'Audio' : type === 'video' ? 'Video' : 'Html';
            const url = `${API_BASE}/packages/${state.currentPackageId}/media/${mediaType}/${encodeURIComponent(item.value)}`;
            const typeLabel = type.charAt(0).toUpperCase() + type.slice(1);
            html += `
                <div class="content-item content-item-media">
                    <div class="content-item-header">
                        <span class="content-item-type">${typeLabel}: ${escapeHtml(item.value)}</span>
                        <button class="btn btn-icon btn-small" onclick="removeContentItem(${ri}, ${ti}, ${qi}, ${i})" title="${t('delete')}">
                            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    ${type === 'image' ? `<img src="${url}" style="max-width:300px;max-height:200px;border-radius:8px;">` : ''}
                    ${type === 'audio' ? `<audio controls src="${url}" style="width:100%;"></audio>` : ''}
                    ${type === 'video' ? `<video controls src="${url}" style="max-width:400px;max-height:300px;"></video>` : ''}
                </div>
            `;
        }
    }
    if (content.length === 0) {
        html += `
            <div class="form-group">
                <label class="form-label">${t('questionText')}</label>
                <textarea class="form-input" rows="4" onchange="addQuestionTextContent(${ri}, ${ti}, ${qi}, this.value)"></textarea>
            </div>
        `;
    }
    html += '</div>';
    // Add content buttons
    html += `
        <div class="add-content-buttons">
            <button class="btn btn-secondary btn-small" onclick="addContentItem(${ri}, ${ti}, ${qi}, 'text')">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
                </svg>
                ${t('addText')}
            </button>
            <div class="dropdown-inline">
                <button class="btn btn-secondary btn-small" onclick="toggleMediaPicker(${ri}, ${ti}, ${qi})">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                        <circle cx="8.5" cy="8.5" r="1.5"/>
                        <polyline points="21 15 16 10 5 21"/>
                    </svg>
                    ${t('addMedia')}
                </button>
                <div id="media-picker-${ri}-${ti}-${qi}" class="media-picker hidden">
                    ${renderMediaPicker(ri, ti, qi)}
                </div>
            </div>
        </div>
    `;
    return html;
}
function renderMediaPicker(ri, ti, qi) {
    const types = [
        { key: 'Images', label: t('images'), type: 'image' },
        { key: 'Audio', label: t('audio'), type: 'audio' },
        { key: 'Video', label: t('video'), type: 'video' }
    ];
    let html = '<div class="media-picker-content">';
    for (const mediaType of types) {
        const files = state.media[mediaType.key] || [];
        if (files.length > 0) {
            html += `<div class="media-picker-section">
                <div class="media-picker-title">${mediaType.label}</div>
                <div class="media-picker-list">`;
            for (const file of files) {
                html += `<div class="media-picker-item" onclick="insertMediaContent(${ri}, ${ti}, ${qi}, '${mediaType.type}', '${escapeHtml(file)}')">
                    ${mediaType.type === 'image' ? `<img src="${API_BASE}/packages/${state.currentPackageId}/media/${mediaType.key}/${encodeURIComponent(file)}" class="media-picker-thumb">` : ''}
                    ${mediaType.type === 'audio' ? `<span class="media-picker-icon">🎵</span>` : ''}
                    ${mediaType.type === 'video' ? `<span class="media-picker-icon">🎬</span>` : ''}
                    <span class="media-picker-name">${escapeHtml(file.length > 20 ? file.substring(0, 18) + '...' : file)}</span>
                </div>`;
            }
            html += '</div></div>';
        }
    }
    if (!state.media.Images?.length && !state.media.Audio?.length && !state.media.Video?.length) {
        html += `<div class="media-picker-empty">${t('noMediaUploaded')}</div>`;
    }
    html += '</div>';
    return html;
}
function toggleMediaPicker(ri, ti, qi) {
    const picker = document.getElementById(`media-picker-${ri}-${ti}-${qi}`);
    if (picker) {
        picker.classList.toggle('hidden');
    }
}
function renderAnswerContent(question, ri, ti, qi) {
    // Handle both contentValue array and simpleValue string for answer parameter
    let content = question.parameters?.answer?.contentValue || [];
    // If no contentValue but has simpleValue, convert it to content item
    if (content.length === 0 && question.parameters?.answer?.simpleValue) {
        content = [{ type: 'text', value: question.parameters.answer.simpleValue.trim() }];
    }
    let html = '<div class="answer-content-list">';
    for (let i = 0; i < content.length; i++) {
        const item = content[i];
        const type = item.type || 'text';
        if (type === 'text') {
            html += `
                <div class="content-item">
                    <div class="content-item-header">
                        <span class="content-item-type">${t('text')}</span>
                        <button class="btn btn-icon btn-small" onclick="removeAnswerContentItem(${ri}, ${ti}, ${qi}, ${i})" title="${t('delete')}">
                            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    <textarea class="form-input" rows="3" onchange="updateAnswerContentItem(${ri}, ${ti}, ${qi}, ${i}, 'value', this.value)">${escapeHtml(item.value || '')}</textarea>
                </div>
            `;
        } else if (item.isRef) {
            // Media reference
            const mediaType = type === 'image' ? 'Images' : type === 'audio' ? 'Audio' : type === 'video' ? 'Video' : 'Html';
            const url = `${API_BASE}/packages/${state.currentPackageId}/media/${mediaType}/${encodeURIComponent(item.value)}`;
            const typeLabel = type.charAt(0).toUpperCase() + type.slice(1);
            html += `
                <div class="content-item content-item-media">
                    <div class="content-item-header">
                        <span class="content-item-type">${typeLabel}: ${escapeHtml(item.value)}</span>
                        <button class="btn btn-icon btn-small" onclick="removeAnswerContentItem(${ri}, ${ti}, ${qi}, ${i})" title="${t('delete')}">
                            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    ${type === 'image' ? `<img src="${url}" style="max-width:300px;max-height:200px;border-radius:8px;">` : ''}
                    ${type === 'audio' ? `<audio controls src="${url}" style="width:100%;"></audio>` : ''}
                    ${type === 'video' ? `<video controls src="${url}" style="max-width:400px;max-height:300px;"></video>` : ''}
                </div>
            `;
        }
    }
    if (content.length === 0) {
        html += `<p class="hint-text" style="color: var(--text-secondary); font-size: 13px; margin: 8px 0;">${t('answerContent')}</p>`;
    }
    html += '</div>';
    // Add content buttons
    html += `
        <div class="add-content-buttons">
            <button class="btn btn-secondary btn-small" onclick="addAnswerContentItem(${ri}, ${ti}, ${qi}, 'text')">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
                </svg>
                ${t('addText')}
            </button>
            <div class="dropdown-inline">
                <button class="btn btn-secondary btn-small" onclick="toggleAnswerMediaPicker(${ri}, ${ti}, ${qi})">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                        <circle cx="8.5" cy="8.5" r="1.5"/>
                        <polyline points="21 15 16 10 5 21"/>
                    </svg>
                    ${t('addMedia')}
                </button>
                <div id="answer-media-picker-${ri}-${ti}-${qi}" class="media-picker hidden">
                    ${renderAnswerMediaPicker(ri, ti, qi)}
                </div>
            </div>
        </div>
    `;
    return html;
}
function renderAnswerMediaPicker(ri, ti, qi) {
    const types = [
        { key: 'Images', label: t('images'), type: 'image' },
        { key: 'Audio', label: t('audio'), type: 'audio' },
        { key: 'Video', label: t('video'), type: 'video' }
    ];
    let html = '<div class="media-picker-content">';
    for (const mediaType of types) {
        const files = state.media[mediaType.key] || [];
        if (files.length > 0) {
            html += `<div class="media-picker-section">
                <div class="media-picker-title">${mediaType.label}</div>
                <div class="media-picker-list">`;
            for (const file of files) {
                html += `<div class="media-picker-item" onclick="insertAnswerMediaContent(${ri}, ${ti}, ${qi}, '${mediaType.type}', '${escapeHtml(file)}')">
                    ${mediaType.type === 'image' ? `<img src="${API_BASE}/packages/${state.currentPackageId}/media/${mediaType.key}/${encodeURIComponent(file)}" class="media-picker-thumb">` : ''}
                    ${mediaType.type === 'audio' ? `<span class="media-picker-icon">🎵</span>` : ''}
                    ${mediaType.type === 'video' ? `<span class="media-picker-icon">🎬</span>` : ''}
                    <span class="media-picker-name">${escapeHtml(file.length > 20 ? file.substring(0, 18) + '...' : file)}</span>
                </div>`;
            }
            html += '</div></div>';
        }
    }
    if (!state.media.Images?.length && !state.media.Audio?.length && !state.media.Video?.length) {
        html += `<div class="media-picker-empty">${t('noMediaUploaded')}</div>`;
    }
    html += '</div>';
    return html;
}
function toggleAnswerMediaPicker(ri, ti, qi) {
    const picker = document.getElementById(`answer-media-picker-${ri}-${ti}-${qi}`);
    if (picker) {
        picker.classList.toggle('hidden');
    }
}
async function addAnswerContentItem(ri, ti, qi, type) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (!question.parameters) question.parameters = {};
    if (!question.parameters.answer) question.parameters.answer = { name: 'answer', contentValue: [] };
    if (!question.parameters.answer.contentValue) question.parameters.answer.contentValue = [];
    question.parameters.answer.contentValue.push({ type: type, value: '' });
    await savePackageToServer();
    renderEditor();
}
async function insertAnswerMediaContent(ri, ti, qi, type, filename) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (!question.parameters) question.parameters = {};
    if (!question.parameters.answer) question.parameters.answer = { name: 'answer', contentValue: [] };
    if (!question.parameters.answer.contentValue) question.parameters.answer.contentValue = [];
    question.parameters.answer.contentValue.push({ type: type, value: filename, isRef: true });
    await savePackageToServer();
    renderEditor();
    showToast(t('mediaAddedToAnswer'));
}
async function removeAnswerContentItem(ri, ti, qi, index) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (question.parameters?.answer?.contentValue) {
        question.parameters.answer.contentValue.splice(index, 1);
        await savePackageToServer();
        renderEditor();
    }
}
async function updateAnswerContentItem(roundIndex, themeIndex, questionIndex, itemIndex, field, value) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    if (question.parameters?.answer?.contentValue?.[itemIndex]) {
        question.parameters.answer.contentValue[itemIndex][field] = value;
        await savePackageToServer();
    }
}
async function addContentItem(ri, ti, qi, type) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (!question.parameters) question.parameters = {};
    if (!question.parameters.question) question.parameters.question = { name: 'question', contentValue: [] };
    if (!question.parameters.question.contentValue) question.parameters.question.contentValue = [];
    question.parameters.question.contentValue.push({ type: type, value: '' });
    await savePackageToServer();
    renderEditor();
}
async function insertMediaContent(ri, ti, qi, type, filename) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (!question.parameters) question.parameters = {};
    if (!question.parameters.question) question.parameters.question = { name: 'question', contentValue: [] };
    if (!question.parameters.question.contentValue) question.parameters.question.contentValue = [];
    question.parameters.question.contentValue.push({ type: type, value: filename, isRef: true });
    await savePackageToServer();
    renderEditor();
    showToast(t('mediaAdded'));
}
async function removeContentItem(ri, ti, qi, index) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (question.parameters?.question?.contentValue) {
        question.parameters.question.contentValue.splice(index, 1);
        await savePackageToServer();
        renderEditor();
    }
}
async function addQuestionTextContent(ri, ti, qi, value) {
    const question = state.currentPackage.rounds[ri].themes[ti].questions[qi];
    if (!question.parameters) question.parameters = {};
    if (!question.parameters.question) question.parameters.question = { name: 'question', contentValue: [] };
    if (!question.parameters.question.contentValue) question.parameters.question.contentValue = [];
    question.parameters.question.contentValue.push({ type: 'text', value: value });
    await savePackageToServer();
}

function getQuestionText(question) {
    if (question.parameters?.question?.contentValue) {
        for (const item of question.parameters.question.contentValue) {
            if (item.type === 'text' || !item.type) {
                return item.value;
            }
        }
    }
    return '';
}

// Update operations
async function updatePackageField(field, value) {
    if (!state.currentPackage) return;
    if (field.includes('.')) {
        const parts = field.split('.');
        let obj = state.currentPackage;
        for (let i = 0; i < parts.length - 1; i++) {
            if (!obj[parts[i]]) obj[parts[i]] = {};
            obj = obj[parts[i]];
        }
        obj[parts[parts.length - 1]] = value;
    } else {
        state.currentPackage[field] = value;
    }
    await savePackageToServer();
    if (field === 'name') {
        const pkg = state.packages.find(p => p.id === state.currentPackageId);
        if (pkg) pkg.name = value;
        renderPackageList();
        renderTreeView();
    }
}

async function updatePackageAuthor(index, value) {
    if (!state.currentPackage.info) state.currentPackage.info = {};
    if (!state.currentPackage.info.authors) state.currentPackage.info.authors = [];
    state.currentPackage.info.authors[index] = value;
    await savePackageToServer();
}

async function addPackageAuthor() {
    if (!state.currentPackage.info) state.currentPackage.info = {};
    if (!state.currentPackage.info.authors) state.currentPackage.info.authors = [];
    state.currentPackage.info.authors.push('');
    await savePackageToServer();
    renderEditor();
}

async function removePackageAuthor(index) {
    state.currentPackage.info.authors.splice(index, 1);
    await savePackageToServer();
    renderEditor();
}

async function updateRound(roundIndex, field, value) {
    state.currentPackage.rounds[roundIndex][field] = value;
    await savePackageToServer();
    if (field === 'name') {
        renderTreeView();
    }
}

async function updateRoundInfo(roundIndex, field, value) {
    if (!state.currentPackage.rounds[roundIndex].info) {
        state.currentPackage.rounds[roundIndex].info = {};
    }
    state.currentPackage.rounds[roundIndex].info[field] = value;
    await savePackageToServer();
}

async function updateTheme(roundIndex, themeIndex, field, value) {
    state.currentPackage.rounds[roundIndex].themes[themeIndex][field] = value;
    await savePackageToServer();
    if (field === 'name') {
        renderTreeView();
    }
}

async function updateThemeInfo(roundIndex, themeIndex, field, value) {
    const theme = state.currentPackage.rounds[roundIndex].themes[themeIndex];
    if (!theme.info) theme.info = {};
    theme.info[field] = value;
    await savePackageToServer();
}

async function updateQuestion(roundIndex, themeIndex, questionIndex, field, value) {
    state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex][field] = value;
    await savePackageToServer();
    if (field === 'price') {
        renderTreeView();
    }
}

async function updateQuestionText(roundIndex, themeIndex, questionIndex, value) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    if (!question.parameters) question.parameters = {};
    question.parameters.question = {
        type: 'content',
        contentValue: [{ type: 'text', value: value }]
    };
    await savePackageToServer();
    renderTreeView();
}

async function updateQuestionContentItem(roundIndex, themeIndex, questionIndex, itemIndex, field, value) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    if (question.parameters?.question?.contentValue?.[itemIndex]) {
        question.parameters.question.contentValue[itemIndex][field] = value;
        await savePackageToServer();
        renderTreeView();
    }
}

async function updateQuestionAnswer(roundIndex, themeIndex, questionIndex, answerType, answerIndex, value) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    question[answerType][answerIndex] = value;
    await savePackageToServer();
}

async function addQuestionAnswer(roundIndex, themeIndex, questionIndex, answerType) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    if (!question[answerType]) question[answerType] = [];
    question[answerType].push('');
    await savePackageToServer();
    renderEditor();
}

async function removeQuestionAnswer(roundIndex, themeIndex, questionIndex, answerType, answerIndex) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    question[answerType].splice(answerIndex, 1);
    await savePackageToServer();
    renderEditor();
}

async function updateQuestionInfo(roundIndex, themeIndex, questionIndex, field, value) {
    const question = state.currentPackage.rounds[roundIndex].themes[themeIndex].questions[questionIndex];
    if (!question.info) question.info = {};
    question.info[field] = value;
    await savePackageToServer();
}

async function savePackageToServer() {
    try {
        await api(`/packages/${state.currentPackageId}`, {
            method: 'PUT',
            body: JSON.stringify(state.currentPackage)
        });
    } catch (err) {
        showToast(t('failedToSave') + ': ' + err.message, 'error');
    }
}

// Add/Delete operations
async function addRound() {
    if (!state.currentPackageId) return;
    try {
        const result = await api(`/packages/${state.currentPackageId}/rounds`, {
            method: 'POST',
            body: JSON.stringify({ name: 'New Round' })
        });
        state.currentPackage.rounds.push(result);
        renderTreeView();
        showToast(t('roundAdded'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function deleteRound(roundIndex) {
    if (!confirm(t('confirmDeleteRound'))) return;
    try {
        await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}`, {
            method: 'DELETE'
        });
        state.currentPackage.rounds.splice(roundIndex, 1);
        state.selectedNode = null;
        renderTreeView();
        renderEditor();
        showToast(t('roundDeleted'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function addTheme(roundIndex) {
    try {
        const result = await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}/themes`, {
            method: 'POST',
            body: JSON.stringify({ name: 'New Theme' })
        });
        state.currentPackage.rounds[roundIndex].themes.push(result);
        renderTreeView();
        showToast(t('themeAdded'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function deleteTheme(roundIndex, themeIndex) {
    if (!confirm(t('confirmDeleteTheme'))) return;
    try {
        await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}/themes/${themeIndex}`, {
            method: 'DELETE'
        });
        state.currentPackage.rounds[roundIndex].themes.splice(themeIndex, 1);
        state.selectedNode = null;
        renderTreeView();
        renderEditor();
        showToast(t('themeDeleted'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function addQuestion(roundIndex, themeIndex) {
    try {
        const result = await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}/themes/${themeIndex}/questions`, {
            method: 'POST',
            body: JSON.stringify({ price: 100, right: [''] })
        });
        state.currentPackage.rounds[roundIndex].themes[themeIndex].questions.push(result);
        renderTreeView();
        showToast(t('questionAdded'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

async function deleteQuestion(roundIndex, themeIndex, questionIndex) {
    if (!confirm(t('confirmDeleteQuestion'))) return;
    try {
        await api(`/packages/${state.currentPackageId}/rounds/${roundIndex}/themes/${themeIndex}/questions/${questionIndex}`, {
            method: 'DELETE'
        });
        state.currentPackage.rounds[roundIndex].themes[themeIndex].questions.splice(questionIndex, 1);
        state.selectedNode = null;
        renderTreeView();
        renderEditor();
        showToast(t('questionDeleted'));
    } catch (err) {
        showToast(err.message, 'error');
    }
}

// Utility
function escapeHtml(text) {
    if (!text) return '';
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Panel Resize functionality
function initResizeHandles() {
    const handles = document.querySelectorAll('.resize-handle');
    handles.forEach(handle => {
        handle.addEventListener('mousedown', startResize);
    });
}

let resizeState = {
    active: false,
    handle: null,
    target: null,
    startX: 0,
    startWidth: 0
};

function startResize(e) {
    e.preventDefault();
    const handle = e.target;
    const targetId = handle.dataset.target;
    const target = document.getElementById(targetId);
    if (!target) return;
    resizeState = {
        active: true,
        handle: handle,
        target: target,
        startX: e.clientX,
        startWidth: target.offsetWidth
    };
    handle.classList.add('active');
    document.body.classList.add('resizing');
    document.addEventListener('mousemove', doResize);
    document.addEventListener('mouseup', stopResize);
}

function doResize(e) {
    if (!resizeState.active) return;
    const diff = e.clientX - resizeState.startX;
    const newWidth = resizeState.startWidth + diff;
    const minWidth = parseInt(getComputedStyle(resizeState.target).minWidth) || 150;
    const maxWidth = parseInt(getComputedStyle(resizeState.target).maxWidth) || 600;
    if (newWidth >= minWidth && newWidth <= maxWidth) {
        resizeState.target.style.width = newWidth + 'px';
        // Save to localStorage
        localStorage.setItem(`panel-width-${resizeState.target.id}`, newWidth);
    }
}

function stopResize() {
    if (resizeState.handle) {
        resizeState.handle.classList.remove('active');
    }
    document.body.classList.remove('resizing');
    document.removeEventListener('mousemove', doResize);
    document.removeEventListener('mouseup', stopResize);
    resizeState.active = false;
}

function restorePanelWidths() {
    const panels = ['sidebar', 'tree-panel'];
    panels.forEach(id => {
        const savedWidth = localStorage.getItem(`panel-width-${id}`);
        if (savedWidth) {
            const panel = document.getElementById(id);
            if (panel) {
                panel.style.width = savedWidth + 'px';
            }
        }
    });
}

// Initialize resize handles on page load
document.addEventListener('DOMContentLoaded', () => {
    initResizeHandles();
    restorePanelWidths();
});
